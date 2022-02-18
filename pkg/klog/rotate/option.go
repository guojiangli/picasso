package rotate

import "picasso/pkg/config"

// Config ...
type RotateOption struct {
	FileName string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Compress   bool
	LocalTime  bool
}

func NewOption() *RotateOption {
	return new(RotateOption)
}

// RawConfig ...
func ConfigOption(key string) (option *RotateOption, err error) {
	option = NewOption()
	err = config.UnmarshalKey(key, option)
	if err != nil {
		return nil, err
	}
	return option, nil
}

func (c *RotateOption) SetFileName(l string) *RotateOption {
	c.FileName = l
	return c
}

func (c *RotateOption) SetMaxSize(w int) *RotateOption {
	c.MaxSize = w
	return c
}

func (c *RotateOption) SetMaxAge(w int) *RotateOption {
	c.MaxAge = w
	return c
}

func (c *RotateOption) SetMaxBackup(w int) *RotateOption {
	c.MaxBackups = w
	return c
}

func (c *RotateOption) SetCompress(w bool) *RotateOption {
	c.Compress = w
	return c
}

func (c *RotateOption) SetLocalTime(w bool) *RotateOption {
	c.LocalTime = w
	return c
}

func SetFileName(l string) *RotateOption {
	c := NewOption()
	c.FileName = l
	return c
}

func SetMaxSize(w int) *RotateOption {
	c := NewOption()
	c.MaxSize = w
	return c
}

func SetMaxAge(w int) *RotateOption {
	c := NewOption()
	c.MaxAge = w
	return c
}

func SetMaxBackup(w int) *RotateOption {
	c := NewOption()
	c.MaxBackups = w
	return c
}

func SetCompress(w bool) *RotateOption {
	c := NewOption()
	c.Compress = w
	return c
}

func SetLocalTime(w bool) *RotateOption {
	c := NewOption()
	c.LocalTime = w
	return c
}

func (c *RotateOption) MergeOption(opts ...*RotateOption) *RotateOption {
	for _, opt := range opts {
		if opt.FileName != "" {
			c.FileName = opt.FileName
		}
		if opt.MaxAge != 0 {
			c.MaxAge = opt.MaxAge
		}
		if opt.MaxSize != 0 {
			c.MaxSize = opt.MaxSize
		}
		if opt.MaxBackups != 0 {
			c.MaxBackups = opt.MaxBackups
		}
		if opt.Compress {
			c.Compress = opt.Compress
		}
		if opt.LocalTime {
			c.LocalTime = opt.LocalTime
		}
	}
	return c
}

func defaultOption() *RotateOption {
	return &RotateOption{
		FileName:   "picasso.log",
		MaxSize:    512,
		MaxAge:     30,
		MaxBackups: 30,
		Compress:   false,
		LocalTime:  true,
	}
}
