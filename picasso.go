package main

import (
	"context"
	"errors"
	"sync"

	"picasso/pkg/klog"
	"picasso/pkg/server"
	"picasso/pkg/signals"
	"picasso/pkg/utils/kcycle"
	"picasso/pkg/utils/kdefer"
	"picasso/pkg/utils/kgo"
	"picasso/pkg/worker"
	"github.com/spf13/pflag"
)

type Stage uint8

const (
	// StageAfterStop after app stop
	StageAfterStop Stage = iota + 1
	// StageBeforeStop before app stop
	StageBeforeStop
)

type BaseConfig struct {
	// 日志路径,为空终端打印
	LogPath string
	// 配置源:nacos/file
	ConfigSource string
	// 配置路径:
	// file: 文件地址
	// nacos: host+port，无http前缀
	ConfigPath string
	// 是否watch
	EnableWatch bool
	// etcd的目录
	PropertyKey string
	// 用户名
	Username string
	// 密码
	Password string
	// nacos data_id
	DataID string
	// nacos group
	Group string
	// nacos config type
	Suffix string
	// nacos namespaceID
	NamespaceID string
}

type Application struct {
	mu          *sync.RWMutex
	cycle       *kcycle.Cycle
	servers     []server.Server
	workers     []worker.Worker
	hooks       map[Stage]*kdefer.DeferStack
	initOnce    sync.Once
	stopOnce    sync.Once
}

var app *Application

var baseConfig BaseConfig

type LoadBaseConfigFunction func()

func init() {
	app = &Application{
		mu:      &sync.RWMutex{},
		cycle:   kcycle.NewCycle(),
		servers: make([]server.Server, 0),
		workers: make([]worker.Worker, 0),
	}
	app.initHooks(StageBeforeStop, StageAfterStop)
}

func LoadWithFlags() LoadBaseConfigFunction {
	return func() {
		pflag.StringVar(&baseConfig.ConfigSource, "source", "", "Input Config Source")
		pflag.StringVar(&baseConfig.ConfigPath, "config", "", "Input Config Path")
		pflag.StringVar(&baseConfig.LogPath, "log", "", "Input Log Path")
		pflag.BoolVar(&baseConfig.EnableWatch, "watch", true, "Enable Watch")
		pflag.StringVar(&baseConfig.Username, "user", "", "Username For Nacos")
		pflag.StringVar(&baseConfig.Password, "pass", "", "Password For Nacos")
		pflag.StringVar(&baseConfig.DataID, "dataid", "", "DataID For Nacos")
		pflag.StringVar(&baseConfig.Group, "group", "", "Group For Nacos")
		pflag.StringVar(&baseConfig.NamespaceID, "namespaceid", "", "Namespaceid For Nacos")
		pflag.StringVar(&baseConfig.Suffix, "suffix", "", "Config Suffix For Nacos")
		pflag.Parse()
	}
}

func LoadWithInput(b *BaseConfig) LoadBaseConfigFunction {
	return func() {
		baseConfig = *b
	}
}

func Init(baseConfig LoadBaseConfigFunction) (err error) {
	app.initOnce.Do(func() {
		baseConfig()
		err = app.initialize()
		if err != nil {
			klog.DefaultLogger().Error().Err(err).Msg("app initialize fail")
			return
		}
	})
	return
}

func (app *Application) initialize() (err error) {
	err = kgo.SerialUntilError(
		app.initConfig,
		app.initDefaultLogger,
	)()
	return
}

func AddServer(s server.Server) {
	app.mu.Lock()
	defer app.mu.Unlock()
	app.servers = append(app.servers, s)
}

func AddWorker(w worker.Worker) {
	app.mu.Lock()
	defer app.mu.Unlock()
	app.workers = append(app.workers, w)
}

// RegisterHooks register a stage Hook
func RegisterHooks(k Stage, fns ...func() error) error {
	app.mu.Lock()
	defer app.mu.Unlock()
	hooks, ok := app.hooks[k]
	if ok {
		hooks.Push(fns...)
		return nil
	}
	return errors.New("hook stage not found")
}

// Run run application
func Run() {
	app.waitSignals() // start signal listen task in goroutine

	// start servers and govern server
	for _, s := range app.servers {
		klog.DefaultLogger().Info().Interface("info", s.Info()).Msg("picasso server run")
		app.cycle.Run(s.Serve)
	}
	// stop workers
	for _, w := range app.workers {
		klog.DefaultLogger().Info().Msg("picasso worker run")
		app.cycle.Run(w.Run)
	}
	// blocking and wait quit
	for err := range app.cycle.Wait() {
		klog.DefaultLogger().Error().Err(err).Msg("picasso receive err")
		GracefulStop(context.TODO())
	}
	klog.Info("picasso over")
}

// Stop application immediately after necessary cleanup
func Stop() (err error) {
	app.stopOnce.Do(func() {
		app.runHooks(StageBeforeStop)
		// stop servers
		app.mu.RLock()
		for _, s := range app.servers {
			klog.DefaultLogger().Info().Interface("info", s.Info()).Msg("picasso server stop")
			app.cycle.Run(s.Stop)
		}
		// stop workers
		for _, w := range app.workers {
			klog.DefaultLogger().Info().Msg("picasso worker stop")
			app.cycle.Run(w.Stop)
		}
		app.mu.RUnlock()
		app.runHooks(StageAfterStop)
		app.cycle.Close()
	})
	return
}

// GracefulStop application after necessary cleanup
func GracefulStop(ctx context.Context) (err error) {
	app.stopOnce.Do(func() {
		app.runHooks(StageBeforeStop)
		// stop servers
		app.mu.RLock()
		for _, s := range app.servers {
			klog.DefaultLogger().Info().Interface("info", s.Info()).Msg("picasso server stop")
			app.cycle.Run(func() error {
				return s.GracefulStop(ctx)
			})
		}
		// stop workers
		for _, w := range app.workers {
			klog.DefaultLogger().Info().Msg("picasso server stop")
			app.cycle.Run(w.Stop)
		}
		app.mu.RUnlock()
		app.runHooks(StageAfterStop)
		app.cycle.Close()
	})
	return
}

// waitSignals wait signal
func (app *Application) waitSignals() {
	signals.Shutdown(func(grace bool) { // when get shutdown signal
		if grace {
			GracefulStop(context.TODO())
		} else {
			Stop()
		}
	})
}

// run hooks
func (app *Application) runHooks(k Stage) {
	hooks, ok := app.hooks[k]
	if ok {
		hooks.Clean()
	}
}
