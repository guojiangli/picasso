package baselogger

import (
	"github.com/alibaba/sentinel-golang/logging"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/nacos-group/nacos-sdk-go/common/logger"
	"github.com/robfig/cron/v3"
	"github.com/uber/jaeger-client-go"
	glogger "gorm.io/gorm/logger"
)

type Logger interface {
	Log(keyvals ...interface{}) error
}

type GormLogger interface {
	glogger.Interface
	Logger
}

type RocketmqLogger interface {
	rlog.Logger
	Logger
}

type KcronLogger interface {
	cron.Logger
	Logger
}

type JaegerLogger interface {
	jaeger.Logger
	Logger
}

type SentinelLogger interface {
	logging.Logger
	Logger
}

type NacosLogger interface {
	logger.Logger
	Logger
	Level(level string)
}
