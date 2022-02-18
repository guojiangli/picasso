package kgin

import (
	"github.com/guojiangli/picasso/pkg/config"
	"github.com/guojiangli/picasso/pkg/klog"
	"github.com/guojiangli/picasso/pkg/klog/baselogger"
	"github.com/gin-gonic/gin"
)

// Config HTTP config
type Option struct {
	Host        string
	Port        int
	Logger      baselogger.Logger
	Middlewares []gin.HandlerFunc
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

func (c *Option) SetHost(s string) *Option {
	c.Host = s
	return c
}

func SetHost(s string) *Option {
	c := NewOption()
	c.Host = s
	return c
}

func SetPort(l int) *Option {
	c := NewOption()
	c.Port = l
	return c
}

func (c *Option) SetPort(l int) *Option {
	c.Port = l
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

func AddMiddleware(l gin.HandlerFunc) *Option {
	c := NewOption()
	c.Middlewares = append(c.Middlewares, l)
	return c
}

func (c *Option) AddMiddleware(l gin.HandlerFunc) *Option {
	c.Middlewares = append(c.Middlewares, l)
	return c
}

func (c *Option) MergeOption(opts ...*Option) *Option {
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if opt.Logger != nil {
			c.Logger = opt.Logger
		}
		if opt.Middlewares != nil {
			c.Middlewares = append(c.Middlewares, opt.Middlewares...)
		}
		if opt.Host != "" {
			c.Host = opt.Host
		}
		if opt.Port != 0 {
			c.Port = opt.Port
		}
	}
	return c
}

func defaultOption() *Option {
	l := klog.DefaultLogger().With().Str("picasso", "kgin").Logger()
	return &Option{
		Logger: klog.NewLoggerWith(&l),
	}
}
