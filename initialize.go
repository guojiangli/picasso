package main

import (
	"io"
	"os"
	"picasso/pkg/config"
	"picasso/pkg/config/source"
	file_source "picasso/pkg/config/source/file"
	"picasso/pkg/klog"
	"picasso/pkg/klog/rotate"
	"picasso/pkg/utils/kdefer"
	"github.com/pkg/errors"
)

func (app *Application) initConfig() (err error) {
	if baseConfig.ConfigPath == "" {
		klog.DefaultLogger().Info().Msg("未找到配置路径,跳过配置初始化")
		return nil
	}
	var provider source.ConfSource
	switch baseConfig.ConfigSource {
	case "file":
		provider, err = file_source.NewSource(&file_source.Option{
			Path:        baseConfig.ConfigPath,
			EnableWatch: baseConfig.EnableWatch,
			Logger:      klog.DefaultLogger(),
		})
		if err != nil {
			klog.DefaultLogger().Error().Err(err).Msg("初始化文件配置失败")
			return err
		}
	default:
		err = errors.New("未找到配置源")
		return err
	}
	err = config.InitConfig(config.NewOption().SetSource(provider))
	if err != nil {
		klog.DefaultLogger().Error().Err(err).Msg("初始化配置出错")
		return err
	}
	klog.DefaultLogger().Info().Msg("kepler完成配置初始化")
	return nil
}

func (app *Application) initDefaultLogger() (err error) {
	var logWriter io.Writer
	if baseConfig.LogPath == "" {
		klog.DefaultLogger().Info().Msg("日志路径为空,日志打印到终端")
		logWriter = os.Stderr
	} else {
		option, err := rotate.ConfigOption("klog.rotate")
		if err != nil {
			klog.DefaultLogger().Error().Err(err).Msg("获取kepler.rotate标准配置失败")
			return err
		}
		logWriter = rotate.NewRotate(option, option.SetFileName(baseConfig.LogPath))
		klog.DefaultLogger().Info().Str("路径", baseConfig.LogPath).Msg("日志文件初始化")
	}
	if baseConfig.ConfigPath == "" && baseConfig.ConfigSource == "" {
		klog.DefaultLogger().Info().Msg("日志配置为空,默认初始化")
		klog.InitLogger()
	} else {
		klog.DefaultLogger().Info().Interface("klog", config.Get("klog")).Msg("获取klog配置")
		option, err := klog.ConfigOption("klog")
		if err != nil {
			klog.DefaultLogger().Error().Err(err).Msg("获取klog标准配置失败")
			return err
		}
		klog.InitLogger(option.SetWriter(logWriter))
		klog.DefaultLogger().AutoLevel("klog")
	}
	return nil
}

// init hooks
func (app *Application) initHooks(hookKeys ...Stage) {
	app.hooks = make(map[Stage]*kdefer.DeferStack, len(hookKeys))
	for _, k := range hookKeys {
		app.hooks[k] = kdefer.NewStack()
	}
}
