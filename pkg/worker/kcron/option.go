package kcron

import (
	"github.com/guojiangli/picasso/pkg/klog"
	"github.com/guojiangli/picasso/pkg/klog/baselogger"
	"github.com/guojiangli/picasso/pkg/worker/kcron/driver"
)

type Option struct {
	WorkerName string
	Driver     driver.Driver
	Logger     baselogger.KcronLogger
}

func NewOption() *Option {
	return new(Option)
}

func (c *Option) SetWorkerName(s string) *Option {
	c.WorkerName = s
	return c
}

func (c *Option) SetDriver(d driver.Driver) *Option {
	c.Driver = d
	return c
}

func (c *Option) SetLogger(l baselogger.KcronLogger) *Option {
	c.Logger = l
	return c
}

func SetWorkerName(s string) *Option {
	c := NewOption()
	c.WorkerName = s
	return c
}

func SetDriver(d driver.Driver) *Option {
	c := NewOption()
	c.Driver = d
	return c
}

func SetLogger(l baselogger.KcronLogger) *Option {
	c := NewOption()
	c.Logger = l
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
		if opt.Driver != nil {
			c.Driver = opt.Driver
		}
		if opt.WorkerName != "" {
			c.WorkerName = opt.WorkerName
		}
	}
	return c
}

func defaultOption() *Option {
	l := klog.DefaultLogger().With().Str("picasso", "kcron").Logger()
	return &Option{
		Logger: &klog.KcronLogger{
			Logger: klog.NewLoggerWith(&l),
		},
	}
}
