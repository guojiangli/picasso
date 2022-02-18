package kredis

import (
	"context"
	"github.com/guojiangli/picasso/pkg/klog/baselogger"
	"github.com/guojiangli/picasso/pkg/utils/ktime"
	"github.com/go-redis/redis/v8"
)

// RedisStub is an unit that handles master and slave.
type Cache struct {
	*redis.Client
	Logger baselogger.Logger
}

func New(opts ...*Option) (*Cache, error) {
	kredisOpts := defaultOption().MergeOption(opts...)
	kredisClient := Cache{
		Logger: kredisOpts.Logger,
	}
	kredisClient.Client = redis.NewClient(&redis.Options{
		Addr:         kredisOpts.Addr,
		Password:     kredisOpts.Password,
		DB:           kredisOpts.DB,
		MaxRetries:   kredisOpts.MaxRetries,
		DialTimeout:  ktime.Duration(kredisOpts.DialTimeout),
		ReadTimeout:  ktime.Duration(kredisOpts.ReadTimeout),
		WriteTimeout: ktime.Duration(kredisOpts.WriteTimeout),
		PoolSize:     kredisOpts.PoolSize,
		MinIdleConns: kredisOpts.MinIdleConns,
		IdleTimeout:  ktime.Duration(kredisOpts.IdleTimeout),
	})
	if err := kredisClient.Ping(context.TODO()).Err(); err != nil {
		kredisClient.Logger.Log("start redis", err)
		return nil, err
	}
	return &kredisClient, nil
}
