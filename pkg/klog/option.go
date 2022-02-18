package klog

import (
	"io"
	"os"

	"picasso/pkg/config"
)

type Option struct {
	Writer io.Writer
	Level  Level
}

func NewOption() *Option {
	return new(Option)
}

func SetLevel(l Level) *Option {
	c := NewOption()
	c.Level = l
	return c
}

func (c *Option) SetLevel(l Level) *Option {
	c.Level = l
	return c
}

func SetWriter(w io.Writer) *Option {
	c := NewOption()
	c.Writer = w
	return c
}

func (c *Option) SetWriter(w io.Writer) *Option {
	c.Writer = w
	return c
}

// RawConfig ...
func ConfigOption(key string) (*Option, error) {
	level := config.GetString(key + ".level")
	lv, err := LevelUnmarshalText(level)
	if err != nil {
		return nil, err
	}
	return NewOption().SetLevel(lv), nil
}

func (c *Option) MergeOption(opts ...*Option) *Option {
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if opt.Level != 0 {
			c.Level = opt.Level
		}
		if opt.Writer != nil {
			c.Writer = opt.Writer
		}
	}
	return c
}

func defaultOption() *Option {
	return &Option{
		Level:  DebugLevel,
		Writer: os.Stderr,
	}
}
