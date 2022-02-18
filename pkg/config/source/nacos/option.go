package nacos

import (
	"github.com/guojiangli/picasso/pkg/klog"
	"github.com/guojiangli/picasso/pkg/klog/baselogger"
)

type Option struct {
	DataID      string
	Group       string
	Suffix      string
	EnableWatch bool
	Logger      baselogger.Logger
}

type Credential struct{}

func NewOption() *Option {
	return new(Option)
}

func SetDataID(s string) *Option {
	c := NewOption()
	c.DataID = s
	return c
}

func (c *Option) SetDataID(s string) *Option {
	c.DataID = s
	return c
}

func SetGroup(s string) *Option {
	c := NewOption()
	c.Group = s
	return c
}

func (c *Option) SetGroup(s string) *Option {
	c.Group = s
	return c
}

func SetSuffix(s string) *Option {
	c := NewOption()
	c.Suffix = s
	return c
}

func (c *Option) SetSuffix(s string) *Option {
	c.Suffix = s
	return c
}

func SetEnableWatch(b bool) *Option {
	c := NewOption()
	c.EnableWatch = b
	return c
}

func (c *Option) SetEnableWatch(b bool) *Option {
	c.EnableWatch = b
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
		if !opt.EnableWatch {
			c.EnableWatch = opt.EnableWatch
		}
		if opt.Logger != nil {
			c.Logger = opt.Logger
		}
		if opt.DataID != "" {
			c.DataID = opt.DataID
		}
		if opt.Group != "" {
			c.Group = opt.Group
		}
		if opt.Suffix != "" {
			c.Suffix = opt.Suffix
		}
	}
	return c
}

func defaultOption() *Option {
	l := klog.DefaultLogger().With().Str("picasso", "source-nacos").Logger()
	return &Option{
		EnableWatch: true,
		Logger:      klog.NewLoggerWith(&l),
	}
}
