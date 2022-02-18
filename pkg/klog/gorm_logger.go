package klog

import (
	"context"
	"time"

	glogger "gorm.io/gorm/logger"
)

type GormLogger struct {
	Logger        *Logger
	logLevel      glogger.LogLevel
	SlowThreshold time.Duration
}

func (w *GormLogger) LogMode(level glogger.LogLevel) glogger.Interface {
	newlogger := *w
	newlogger.logLevel = level
	return &newlogger
}

func (w GormLogger) Info(_ context.Context, msg string, data ...interface{}) {
	if w.logLevel >= glogger.Info {
		w.Logger.Info().Interface("details", data).Msg(msg)
	}
}

func (w GormLogger) Warn(_ context.Context, msg string, data ...interface{}) {
	if w.logLevel >= glogger.Warn {
		w.Logger.Warn().Interface("details", data).Msg(msg)
	}
}

func (w GormLogger) Error(_ context.Context, msg string, data ...interface{}) {
	if w.logLevel >= glogger.Error {
		w.Logger.Error().Interface("details", data).Msg(msg)
	}
}

func (w GormLogger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	if w.logLevel > glogger.Silent {
		elapsed := time.Since(begin)
		switch {
		case err != nil && w.logLevel >= glogger.Error:
			sql, rows := fc()
			w.Logger.Error().Str("sql", sql).Int64("rows", rows).Float64("elapsed", float64(elapsed.Nanoseconds())/1e6).Err(err).Send()
		case elapsed > w.SlowThreshold && w.SlowThreshold != 0 && w.logLevel >= glogger.Warn:
			sql, rows := fc()
			w.Logger.Warn().Str("sql", sql).Int64("rows", rows).Float64("elapsed", float64(elapsed.Nanoseconds())/1e6).Send()
		case w.logLevel >= glogger.Info:
			sql, rows := fc()
			w.Logger.Info().Str("sql", sql).Int64("rows", rows).Float64("elapsed", float64(elapsed.Nanoseconds())/1e6).Send()
		}
	}
}

func (w *GormLogger) Log(keyvals ...interface{}) error {
	w.Logger.Info().Interface("details", keyvals).Send()
	return nil
}
