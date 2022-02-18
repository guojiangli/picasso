package kgorm

import (
	"picasso/pkg/klog/baselogger"
	"picasso/pkg/utils/ktime"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

type DB struct {
	*gorm.DB
	InstanceName string
	Logger       baselogger.Logger
}

func New(dialector gorm.Dialector, opts ...*Option) (*DB, error) {
	gormOpts := defaultOption().MergeOption(opts...)
	DB := DB{
		Logger:       gormOpts.Logger,
		InstanceName: gormOpts.InstanceName,
	}
	db, err := gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction: gormOpts.SkipDefaultTransaction,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   gormOpts.TablePrefix,   // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: gormOpts.SingularTable, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
		Logger:                                   gormOpts.Logger,
		DryRun:                                   gormOpts.DryRun,
		PrepareStmt:                              gormOpts.PrepareStmt,
		DisableAutomaticPing:                     gormOpts.DisableAutomaticPing,
		DisableForeignKeyConstraintWhenMigrating: gormOpts.DisableForeignKeyConstraintWhenMigrating,
	})
	if err != nil {
		DB.Logger.Log("创建GORM连接池失败", err)
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		DB.Logger.Log("获取GORM的sqlDB失败", err)
		return nil, err
	}
	sqlDB.SetConnMaxLifetime(ktime.Duration(gormOpts.ConnMaxLifetime))
	sqlDB.SetMaxIdleConns(gormOpts.MaxIdleConns)
	sqlDB.SetMaxOpenConns(gormOpts.MaxOpenConns)
	DB.DB = db
	_instances.Store(DB.InstanceName, &DB)
	return &DB, nil
}

func (db *DB) RegistResolver(sources []gorm.Dialector, replicas []gorm.Dialector, datas ...interface{}) (*DB, error) {
	config := dbresolver.Config{}
	for _, s := range sources {
		config.Sources = append(config.Sources, s)
	}
	for _, r := range replicas {
		config.Replicas = append(config.Replicas, r)
	}
	config.Policy = dbresolver.RandomPolicy{}
	err := db.Use(dbresolver.Register(config, datas...))
	if err != nil {
		db.Logger.Log("RegistResolver err:", err)
		return nil, err
	}
	return db, nil
}
