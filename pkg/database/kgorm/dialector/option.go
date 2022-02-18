package dialector

import (
	"picasso/pkg/config"
)

type Option struct {
	DSN string
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

func SetDSN(l string) *Option {
	c := NewOption()
	c.DSN = l
	return c
}

func (c *Option) SetDSN(l string) *Option {
	c.DSN = l
	return c
}

func (c *Option) MergeOption(opts ...*Option) *Option {
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if opt.DSN != "" {
			c.DSN = opt.DSN
		}
	}
	return c
}

func defaultOption() *Option {
	return &Option{}
}
