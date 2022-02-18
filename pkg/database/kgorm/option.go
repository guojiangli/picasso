package kgorm

import (
	"picasso/pkg/config"
	"picasso/pkg/klog"
	"picasso/pkg/klog/baselogger"
)

type Option struct {
	InstanceName string
	// 最大空闲连接数
	MaxIdleConns int
	// 最大活动连接数
	MaxOpenConns int
	// 连接的最大存活时间
	ConnMaxLifetime string
	// 生成 SQL 但不执行，可以用于准备或测试生成的 SQL
	DryRun bool
	// PreparedStmt 在执行任何 SQL 时都会创建一个 prepared statement 并将其缓存，以提高后续的效率，
	PrepareStmt bool
	// 在完成初始化后，GORM 会自动 ping 数据库以检查数据库的可用性，若要禁用该特性，可将其设置为 true
	DisableAutomaticPing bool
	// 在 AutoMigrate 或 CreateTable 时，GORM 会自动创建外键约束，若要禁用该特性，可将其设置为 true
	DisableForeignKeyConstraintWhenMigrating bool
	// 慢日志阈值
	SlowThreshold          string
	SkipDefaultTransaction bool
	Logger                 baselogger.GormLogger
	TablePrefix            string
	SingularTable          bool
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

func SetInstanceName(l string) *Option {
	c := NewOption()
	c.InstanceName = l
	return c
}

func (c *Option) SetInstanceName(l string) *Option {
	c.InstanceName = l
	return c
}

func SetMaxIdleConns(l int) *Option {
	c := NewOption()
	c.MaxIdleConns = l
	return c
}

func (c *Option) SetMaxIdleConns(l int) *Option {
	c.MaxIdleConns = l
	return c
}

func SetMaxOpenConns(l int) *Option {
	c := NewOption()
	c.MaxOpenConns = l
	return c
}

func (c *Option) SetMaxOpenConns(l int) *Option {
	c.MaxOpenConns = l
	return c
}

func SetTablePrefix(l string) *Option {
	c := NewOption()
	c.TablePrefix = l
	return c
}

func (c *Option) SetTablePrefix(l string) *Option {
	c.TablePrefix = l
	return c
}

func SetConnMaxLifetime(l string) *Option {
	c := NewOption()
	c.ConnMaxLifetime = l
	return c
}

func (c *Option) SetConnMaxLifetime(l string) *Option {
	c.ConnMaxLifetime = l
	return c
}

func SetDryRun(l bool) *Option {
	c := NewOption()
	c.DryRun = l
	return c
}

func (c *Option) SetDryRun(l bool) *Option {
	c.DryRun = l
	return c
}

func SetPrepareStmt(l bool) *Option {
	c := NewOption()
	c.PrepareStmt = l
	return c
}

func (c *Option) SetPrepareStmt(l bool) *Option {
	c.PrepareStmt = l
	return c
}

func SetDisableAutomaticPing(l bool) *Option {
	c := NewOption()
	c.DisableAutomaticPing = l
	return c
}

func (c *Option) SetDisableAutomaticPing(l bool) *Option {
	c.DisableAutomaticPing = l
	return c
}

func SetDisableForeignKeyConstraintWhenMigrating(l bool) *Option {
	c := NewOption()
	c.DisableForeignKeyConstraintWhenMigrating = l
	return c
}

func (c *Option) SetDisableForeignKeyConstraintWhenMigrating(l bool) *Option {
	c.DisableForeignKeyConstraintWhenMigrating = l
	return c
}

func SetSingularTable(l bool) *Option {
	c := NewOption()
	c.SingularTable = l
	return c
}

func (c *Option) SetSingularTable(l bool) *Option {
	c.SingularTable = l
	return c
}

func SetSkipDefaultTransaction(l bool) *Option {
	c := NewOption()
	c.SkipDefaultTransaction = l
	return c
}

func (c *Option) SetSkipDefaultTransaction(l bool) *Option {
	c.SkipDefaultTransaction = l
	return c
}

func SetLogger(l baselogger.GormLogger) *Option {
	c := NewOption()
	c.Logger = l
	return c
}

func (c *Option) SetLogger(l baselogger.GormLogger) *Option {
	c.Logger = l
	return c
}

func SetSlowThreshold(l string) *Option {
	c := NewOption()
	c.SlowThreshold = l
	return c
}

func (c *Option) SetSlowThreshold(l string) *Option {
	c.SlowThreshold = l
	return c
}

func (c *Option) MergeOption(opts ...*Option) *Option {
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if opt.MaxIdleConns != 0 {
			c.MaxIdleConns = opt.MaxIdleConns
		}
		if opt.MaxOpenConns != 0 {
			c.MaxOpenConns = opt.MaxOpenConns
		}
		if opt.ConnMaxLifetime != "" {
			c.ConnMaxLifetime = opt.ConnMaxLifetime
		}
		if opt.InstanceName != "" {
			c.InstanceName = opt.InstanceName
		}
		if opt.Logger != nil {
			c.Logger = opt.Logger
		}
		if opt.PrepareStmt != false {
			c.PrepareStmt = opt.PrepareStmt
		}
		if opt.DryRun != false {
			c.DryRun = opt.DryRun
		}
		if opt.DisableAutomaticPing != false {
			c.DisableAutomaticPing = opt.DisableAutomaticPing
		}
		if opt.DisableForeignKeyConstraintWhenMigrating != false {
			c.DisableForeignKeyConstraintWhenMigrating = opt.DisableForeignKeyConstraintWhenMigrating
		}
		if opt.TablePrefix != "" {
			c.TablePrefix = opt.TablePrefix
		}
		if opt.SingularTable != false {
			c.SingularTable = opt.SingularTable
		}
	}
	return c
}

func defaultOption() *Option {
	l := klog.DefaultLogger().With().Str("picasso", "gorm").Logger()
	return &Option{
		InstanceName:                             "default",
		MaxIdleConns:                             10,
		MaxOpenConns:                             100,
		ConnMaxLifetime:                          "300s",
		DryRun:                                   false,
		PrepareStmt:                              false,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: true,
		SlowThreshold:                            "1s",
		Logger: &klog.GormLogger{
			Logger: klog.NewLoggerWith(&l),
		},
	}
}
