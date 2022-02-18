package config

import (
	"reflect"
	"strings"
	"sync"

	"picasso/pkg/klog/baselogger"

	"picasso/pkg/config/source"
	"picasso/pkg/utils/kmap"
	"github.com/pkg/errors"
)

const (
	defaultKeyDelim = "."
)

var defaultConfig *Config

// Config provides Config for application.
type Config struct {
	logger    baselogger.Logger
	source    source.ConfSource
	override  map[string]interface{}
	keyMap    *sync.Map
	keyDelim  string
	onChanges []func(*Config)
	mu        sync.RWMutex
}

// New constructs a new Config with provider.
func InitConfig(opts ...*Option) (err error) {
	confOpt := defaultOption().MergeOption(opts...)
	if confOpt.Source == nil {
		err = errors.New("Source not empty")
		return err
	}
	defaultConfig = &Config{
		override:  make(map[string]interface{}),
		keyDelim:  defaultKeyDelim,
		keyMap:    &sync.Map{},
		onChanges: make([]func(*Config), 0),
	}
	defaultConfig.source = confOpt.Source
	err = defaultConfig.loadFromConfSource()
	if err != nil {
		return err
	}
	return nil
}

// LoadFromConfSource 读取配置信息,加载到内存,
// 启动一个协程监听配置变化,同步配置到内存
func (c *Config) loadFromConfSource() error {
	content, err := c.source.ReadConfig()
	if err != nil {
		return err
	}
	if err := c.load(content); err != nil {
		return err
	}
	go func() {
		for con := range c.source.ConfigChanged() {
			c.load(con)
			for _, change := range c.onChanges {
				change(c)
			}
		}
	}()
	return nil
}

// OnChange 注册change回调函数
func OnChange(fn func(*Config)) {
	defaultConfig.onChanges = append(defaultConfig.onChanges, fn)
}

// Load ...
func (c *Config) load(content map[string]interface{}) error {
	return c.apply(content)
}

func (c *Config) apply(conf map[string]interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	changes := make(map[string]interface{})
	kmap.MergeStringMap(c.override, conf)
	for k, v := range c.traverse(c.keyDelim) {
		orig, ok := c.keyMap.Load(k)
		if ok && !reflect.DeepEqual(orig, v) {
			changes[k] = v
		}
		c.keyMap.Store(k, v)
	}
	return nil
}

func SetOverride(key string, val interface{}) error {
	paths := strings.Split(key, defaultConfig.keyDelim)
	lastKey := paths[len(paths)-1]
	m := deepSearch(defaultConfig.override, paths[:len(paths)-1])
	m[lastKey] = val
	return defaultConfig.apply(m)
}
