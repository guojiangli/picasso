package consumer

import (
	"picasso/pkg/config"
	"picasso/pkg/klog"
	"picasso/pkg/klog/baselogger"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

type PushConsumerOption struct {
	GroupName    string
	InstanceName string // 访问多个rocketmq集群才需要设置
	NameServer   string // TCP接入点
	AccessKey    string
	SecretKey    string
	Namespace    string
	// 	BroadCasting
	// 	Clustering
	ConsumerModel string
	// 	lastoffset
	// 	firstoffset
	// 	timestamp
	ConsumeFromWhere           string
	ConsumeOrderly             bool
	ConsumeMessageBatchMaxSize int
	PullBatchSize              int
	AllocateStrategy           consumer.AllocateStrategy
	RetryTimes                 int
	Interceptors               []primitive.Interceptor
	AutoCommit                 bool
	Logger                     baselogger.RocketmqLogger
	LogLevel                   string
}

func NewPushConsumerOption() *PushConsumerOption {
	return new(PushConsumerOption)
}

// RawConfig ...
func ConfigPushConsumerOption(key string) (option *PushConsumerOption, err error) {
	option = NewPushConsumerOption()
	err = config.UnmarshalKey(key, option)
	if err != nil {
		return nil, err
	}
	return option, nil
}

func (c *PushConsumerOption) SetGroupName(l string) *PushConsumerOption {
	c.GroupName = l
	return c
}

func (c *PushConsumerOption) SetNameServer(l string) *PushConsumerOption {
	c.NameServer = l
	return c
}

func (c *PushConsumerOption) SetLogger(l baselogger.RocketmqLogger) *PushConsumerOption {
	c.Logger = l
	return c
}

func (c *PushConsumerOption) SetInstanceName(l string) *PushConsumerOption {
	c.InstanceName = l
	return c
}

func (c *PushConsumerOption) SetAccessKey(l string) *PushConsumerOption {
	c.AccessKey = l
	return c
}

func (c *PushConsumerOption) SetSecretKey(l string) *PushConsumerOption {
	c.SecretKey = l
	return c
}

func (c *PushConsumerOption) SetConsumerModel(l string) *PushConsumerOption {
	c.ConsumerModel = l
	return c
}

func (c *PushConsumerOption) SetConsumeFromWhere(l string) *PushConsumerOption {
	c.ConsumeFromWhere = l
	return c
}

func (c *PushConsumerOption) SetNamespace(l string) *PushConsumerOption {
	c.Namespace = l
	return c
}

func (c *PushConsumerOption) SetConsumeOrderly(l bool) *PushConsumerOption {
	c.ConsumeOrderly = l
	return c
}

func (c *PushConsumerOption) SetConsumeMessageBatchMaxSize(l int) *PushConsumerOption {
	c.ConsumeMessageBatchMaxSize = l
	return c
}

func (c *PushConsumerOption) SetPullBatchSize(l int) *PushConsumerOption {
	c.PullBatchSize = l
	return c
}

func (c *PushConsumerOption) AddInterceptor(f ...primitive.Interceptor) *PushConsumerOption {
	c.Interceptors = append(c.Interceptors, f...)
	return c
}

func (c *PushConsumerOption) SetAllocateStrategy(a consumer.AllocateStrategy) *PushConsumerOption {
	c.AllocateStrategy = a
	return c
}

func (c *PushConsumerOption) SetRetryTimes(l int) *PushConsumerOption {
	c.RetryTimes = l
	return c
}

func (c *PushConsumerOption) SetAutoCommit(l bool) *PushConsumerOption {
	c.AutoCommit = l
	return c
}

func (c *PushConsumerOption) PushConsumerSetLogLevel(l string) *PushConsumerOption {
	c.LogLevel = l
	return c
}

func PushConsumerSetGroupName(l string) *PushConsumerOption {
	c := NewPushConsumerOption()
	c.GroupName = l
	return c
}

func PushConsumerSetLogLevel(l string) *PushConsumerOption {
	c := NewPushConsumerOption()
	c.LogLevel = l
	return c
}

func PushConsumerSetNameServer(l string) *PushConsumerOption {
	c := NewPushConsumerOption()
	c.NameServer = l
	return c
}

func PushConsumerSetLogger(l baselogger.RocketmqLogger) *PushConsumerOption {
	c := NewPushConsumerOption()
	c.Logger = l
	return c
}

func PushConsumerSetInstanceName(l string) *PushConsumerOption {
	c := NewPushConsumerOption()
	c.InstanceName = l
	return c
}

func PushConsumerSetAccessKey(l string) *PushConsumerOption {
	c := NewPushConsumerOption()
	c.AccessKey = l
	return c
}

func PushConsumerSetSecretKey(l string) *PushConsumerOption {
	c := NewPushConsumerOption()
	c.SecretKey = l
	return c
}

func PushConsumerSetConsumerModel(l string) *PushConsumerOption {
	c := NewPushConsumerOption()
	c.ConsumerModel = l
	return c
}

func PushConsumerSetConsumeFromWhere(l string) *PushConsumerOption {
	c := NewPushConsumerOption()
	c.ConsumeFromWhere = l
	return c
}

func PushConsumerSetNamespace(l string) *PushConsumerOption {
	c := NewPushConsumerOption()
	c.Namespace = l
	return c
}

func PushConsumerSetConsumeOrderly(l bool) *PushConsumerOption {
	c := NewPushConsumerOption()
	c.ConsumeOrderly = l
	return c
}

func PushConsumerSetConsumeMessageBatchMaxSize(l int) *PushConsumerOption {
	c := NewPushConsumerOption()
	c.ConsumeMessageBatchMaxSize = l
	return c
}

func PushConsumerSetPullBatchSize(l int) *PushConsumerOption {
	c := NewPushConsumerOption()
	c.PullBatchSize = l
	return c
}

func PushConsumerAddInterceptor(f ...primitive.Interceptor) *PushConsumerOption {
	c := NewPushConsumerOption()
	c.Interceptors = append(c.Interceptors, f...)
	return c
}

func PushConsumerSetAllocateStrategy(a consumer.AllocateStrategy) *PushConsumerOption {
	c := NewPushConsumerOption()
	c.AllocateStrategy = a
	return c
}

func PushConsumerSetRetryTimes(l int) *PushConsumerOption {
	c := NewPushConsumerOption()
	c.RetryTimes = l
	return c
}

func PushConsumerSetAutoCommit(l bool) *PushConsumerOption {
	c := NewPushConsumerOption()
	c.AutoCommit = l
	return c
}

func (c *PushConsumerOption) MergeOption(opts ...*PushConsumerOption) *PushConsumerOption {
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if opt.GroupName != "" {
			c.GroupName = opt.GroupName
		}
		if opt.InstanceName != "" {
			c.InstanceName = opt.InstanceName
		}
		if opt.Namespace != "" {
			c.Namespace = opt.Namespace
		}
		if opt.NameServer != "" {
			c.NameServer = opt.NameServer
		}
		if opt.AccessKey != "" {
			c.AccessKey = opt.AccessKey
		}
		if opt.SecretKey != "" {
			c.SecretKey = opt.SecretKey
		}
		if opt.ConsumerModel != "" {
			c.ConsumerModel = opt.ConsumerModel
		}
		if opt.ConsumeFromWhere != "" {
			c.ConsumeFromWhere = opt.ConsumeFromWhere
		}
		if opt.ConsumeOrderly {
			c.ConsumeOrderly = opt.ConsumeOrderly
		}
		if opt.ConsumeMessageBatchMaxSize != 0 {
			c.ConsumeMessageBatchMaxSize = opt.ConsumeMessageBatchMaxSize
		}
		if opt.PullBatchSize != 0 {
			c.PullBatchSize = opt.PullBatchSize
		}
		if opt.AllocateStrategy != nil {
			c.AllocateStrategy = opt.AllocateStrategy
		}
		if opt.RetryTimes != 0 {
			c.RetryTimes = opt.RetryTimes
		}
		if opt.Interceptors != nil {
			c.Interceptors = opt.Interceptors
		}
		if opt.Logger != nil {
			c.Logger = opt.Logger
		}
		if opt.AutoCommit {
			c.AutoCommit = opt.AutoCommit
		}
		if opt.Logger != nil {
			c.Logger = opt.Logger
		}
		if opt.LogLevel != "" {
			c.LogLevel = opt.LogLevel
		}
	}
	return c
}

func defaultPushConsumerOption() *PushConsumerOption {
	l := klog.DefaultLogger().With().Str("picasso", "rocketmq_consumer").Logger()
	return &PushConsumerOption{
		Logger: &klog.RocketmqLogger{
			Logger: klog.NewLoggerWith(&l),
		},
		LogLevel: "warn",
	}
}
