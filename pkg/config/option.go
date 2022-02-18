package config

import (
	"picasso/pkg/config/source"
)

type Option struct {
	Source source.ConfSource
}

var defaultGetOptionTag = "mapstructure"

func NewOption() *Option {
	return new(Option)
}

func SetSource(s source.ConfSource) *Option {
	c := NewOption()
	c.Source = s
	return c
}

func (c *Option) SetSource(s source.ConfSource) *Option {
	c.Source = s
	return c
}

func (c *Option) MergeOption(opts ...*Option) *Option {
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if opt.Source != nil {
			c.Source = opt.Source
		}
	}
	return c
}

func defaultOption() *Option {
	return &Option{}
}
