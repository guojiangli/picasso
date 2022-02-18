package kredis

import (
	"picasso/pkg/config"
	"picasso/pkg/klog"
	"picasso/pkg/klog/baselogger"
)

type Option struct {
	// Addr 节点连接地址
	Addr string
	// Password 密码
	Password string
	// DB，默认为0, 一般应用不推荐使用DB分片
	DB int
	// PoolSize 集群内每个节点的最大连接池限制 默认每个CPU10个连接
	PoolSize int
	// MaxRedirects 网络相关的错误最大重试次数 默认8次
	MaxRetries int
	// MinIdleConns 最小空闲连接数
	MinIdleConns int
	// DialTimeout 拨超时时间
	DialTimeout string
	// ReadTimeout 读超时 默认3s
	ReadTimeout string
	// WriteTimeout 读超时 默认3s
	WriteTimeout string
	// IdleTimeout 连接最大空闲时间，默认60s, 超过该时间，连接会被主动关闭
	IdleTimeout string
	Logger      baselogger.Logger
}

func NewOption() *Option {
	return new(Option)
}

// RawConfig ...
func ConfigOption(key string) (option *Option, err error) {
	option = NewOption()
	err = config.UnmarshalKey(key, option)
	if err != nil {
		return nil, err
	}
	return option, nil
}

func SetAddr(l string) *Option {
	c := NewOption()
	c.Addr = l
	return c
}

func (c *Option) SetAddr(l string) *Option {
	c.Addr = l
	return c
}

func SetPassword(l string) *Option {
	c := NewOption()
	c.Password = l
	return c
}

func (c *Option) SetPassword(l string) *Option {
	c.Password = l
	return c
}

func SetDB(l int) *Option {
	c := NewOption()
	c.DB = l
	return c
}

func (c *Option) SetDB(l int) *Option {
	c.DB = l
	return c
}

func SetPoolSize(l int) *Option {
	c := NewOption()
	c.PoolSize = l
	return c
}

func (c *Option) SetPoolSize(l int) *Option {
	c.PoolSize = l
	return c
}

func SetMaxRetries(l int) *Option {
	c := NewOption()
	c.MaxRetries = l
	return c
}

func (c *Option) SetMaxRetries(l int) *Option {
	c.MaxRetries = l
	return c
}

func SetMinIdleConns(l int) *Option {
	c := NewOption()
	c.MinIdleConns = l
	return c
}

func (c *Option) SetMinIdleConns(l int) *Option {
	c.MinIdleConns = l
	return c
}

func SetDialTimeout(l string) *Option {
	c := NewOption()
	c.DialTimeout = l
	return c
}

func (c *Option) SetDialTimeout(l string) *Option {
	c.DialTimeout = l
	return c
}

func SetReadTimeout(l string) *Option {
	c := NewOption()
	c.ReadTimeout = l
	return c
}

func (c *Option) SetReadTimeout(l string) *Option {
	c.ReadTimeout = l
	return c
}

func SetWriteTimeout(l string) *Option {
	c := NewOption()
	c.WriteTimeout = l
	return c
}

func (c *Option) SetWriteTimeout(l string) *Option {
	c.WriteTimeout = l
	return c
}

func SetIdleTimeout(l string) *Option {
	c := NewOption()
	c.IdleTimeout = l
	return c
}

func (c *Option) SetIdleTimeout(l string) *Option {
	c.IdleTimeout = l
	return c
}

func SetLogger(l baselogger.Logger) *Option {
	c := NewOption()
	c.Logger = l
	return c
}

func (c *Option) SetLogger(l baselogger.Logger) *Option {
	c.Logger = l
	return c
}

func (c *Option) MergeOption(opts ...*Option) *Option {
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if opt.Addr != "" {
			c.Addr = opt.Addr
		}
		if opt.Password != "" {
			c.Password = opt.Password
		}
		if opt.DB != 0 {
			c.DB = opt.DB
		}
		if opt.PoolSize != 0 {
			c.PoolSize = opt.PoolSize
		}
		if opt.MaxRetries != 0 {
			c.MaxRetries = opt.MaxRetries
		}
		if opt.MinIdleConns != 0 {
			c.MinIdleConns = opt.MinIdleConns
		}
		if opt.DialTimeout != "" {
			c.DialTimeout = opt.DialTimeout
		}
		if opt.ReadTimeout != "" {
			c.ReadTimeout = opt.ReadTimeout
		}
		if opt.WriteTimeout != "" {
			c.WriteTimeout = opt.WriteTimeout
		}
		if opt.IdleTimeout != "" {
			c.IdleTimeout = opt.IdleTimeout
		}
		if opt.Logger != nil {
			c.Logger = opt.Logger
		}
	}
	return c
}

func defaultOption() *Option {
	l := klog.DefaultLogger().With().Str("picasso", "redis").Logger()
	return &Option{
		DB:           0,
		PoolSize:     30,
		MaxRetries:   3,
		MinIdleConns: 20,
		DialTimeout:  "1s",
		ReadTimeout:  "3s",
		WriteTimeout: "3s",
		IdleTimeout:  "60s",
		Logger:       klog.NewLoggerWith(&l),
	}
}
