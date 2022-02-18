package consumer

import (
	"picasso/pkg/klog/baselogger"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/rlog"
)

type PushConsumer struct {
	rocketmq.PushConsumer
	Logger baselogger.RocketmqLogger
}

func NewPushConsumer(opts ...*PushConsumerOption) (*PushConsumer, error) {
	opt := defaultPushConsumerOption().MergeOption(opts...)

	var option []consumer.Option

	if opt.GroupName != "" {
		option = append(option, consumer.WithGroupName(opt.GroupName))
	}
	if opt.InstanceName != "" {
		option = append(option, consumer.WithInstance(opt.InstanceName))
	}
	if opt.Namespace != "" {
		option = append(option, consumer.WithNamespace(opt.Namespace))
	}
	if opt.NameServer != "" {
		option = append(option, consumer.WithNameServer([]string{opt.NameServer}))
	}
	if opt.AccessKey != "" || opt.SecretKey != "" {
		option = append(option, consumer.WithCredentials(primitive.Credentials{
			AccessKey: opt.AccessKey,
			SecretKey: opt.SecretKey,
		}))
	}
	if opt.ConsumerModel != "" {
		switch opt.ConsumerModel {
		case "BroadCasting":
			option = append(option, consumer.WithConsumerModel(consumer.BroadCasting))
		case "Clustering":
			option = append(option, consumer.WithConsumerModel(consumer.Clustering))
		}
	}
	if opt.ConsumeFromWhere != "" {
		switch opt.ConsumeFromWhere {
		case "lastOffset":
			option = append(option, consumer.WithConsumeFromWhere((consumer.ConsumeFromWhere)(consumer.ConsumeFromLastOffset)))
		case "firstOffset":
			option = append(option, consumer.WithConsumeFromWhere((consumer.ConsumeFromWhere)(consumer.ConsumeFromFirstOffset)))
		case "timestamp":
			option = append(option, consumer.WithConsumeFromWhere((consumer.ConsumeFromWhere)(consumer.ConsumeFromTimestamp)))
		}
	}

	if opt.ConsumeOrderly != false {
		option = append(option, consumer.WithConsumerOrder(opt.ConsumeOrderly))
	}

	if opt.ConsumeMessageBatchMaxSize != 0 {
		option = append(option, consumer.WithConsumeMessageBatchMaxSize(opt.ConsumeMessageBatchMaxSize))
	}
	if opt.PullBatchSize != 0 {
		option = append(option, consumer.WithPullBatchSize(int32(opt.PullBatchSize)))
	}
	if opt.AllocateStrategy != nil {
		option = append(option, consumer.WithStrategy(opt.AllocateStrategy))
	}

	if opt.AutoCommit != false {
		option = append(option, consumer.WithAutoCommit(opt.AutoCommit))
	}

	if opt.RetryTimes != 0 {
		option = append(option, consumer.WithRetry(opt.RetryTimes))
	}
	if opt.Interceptors != nil {
		option = append(option, consumer.WithInterceptor(opt.Interceptors...))
	}
	consumer, err := rocketmq.NewPushConsumer(option...)
	if err != nil {
		return nil, err
	}
	opt.Logger.Level(opt.LogLevel)
	rlog.SetLogger(opt.Logger)
	return &PushConsumer{
		PushConsumer: consumer,
		Logger:       opt.Logger,
	}, nil
}
