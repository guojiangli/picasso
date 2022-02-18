package file

import (
	"github.com/guojiangli/picasso/pkg/klog"
	"github.com/guojiangli/picasso/pkg/klog/baselogger"
)

type Option struct {
	Path        string
	EnableWatch bool
	Logger      baselogger.Logger
}

func NewOption() *Option {
	return new(Option)
}

func SetPath(s string) *Option {
	c := NewOption()
	c.Path = s
	return c
}

func (c *Option) SetPath(s string) *Option {
	c.Path = s
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
		if opt.Path != "" {
			c.Path = opt.Path
		}
		if !opt.EnableWatch {
			c.EnableWatch = opt.EnableWatch
		}
		if opt.Logger != nil {
			c.Logger = opt.Logger
		}
	}
	return c
}

func defaultOption() *Option {
	l := klog.DefaultLogger().With().Str("picasso", "source-file").Logger()
	return &Option{
		EnableWatch: true,
		Logger:      klog.NewLoggerWith(&l),
	}
}
