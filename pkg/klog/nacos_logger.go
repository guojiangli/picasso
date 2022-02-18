package klog

import (
	"strings"

	"github.com/rs/zerolog"
)

type NacosLogger struct {
	Logger *Logger
}

func (l *NacosLogger) Info(args ...interface{}) {
	l.Logger.Info().Interface("details", args).Send()
}

func (l *NacosLogger) Warn(args ...interface{}) {
	l.Logger.Warn().Interface("details", args).Send()
}

func (l *NacosLogger) Error(args ...interface{}) {
	l.Logger.Error().Interface("details", args).Send()
}

func (l *NacosLogger) Debug(args ...interface{}) {
	l.Logger.Debug().Interface("details", args).Send()
}

func (l *NacosLogger) Infof(fmt string, args ...interface{}) {
	l.Logger.Info().Msgf(fmt, args)
}

func (l *NacosLogger) Warnf(fmt string, args ...interface{}) {
	l.Logger.Warn().Msgf(fmt, args)
}

func (l *NacosLogger) Errorf(fmt string, args ...interface{}) {
	l.Logger.Error().Msgf(fmt, args)
}

func (l *NacosLogger) Debugf(fmt string, args ...interface{}) {
	l.Logger.Debug().Msgf(fmt, args)
}

func (l *NacosLogger) Log(keyvals ...interface{}) error {
	l.Logger.Info().Interface("details", keyvals)
	return nil
}

func (l *NacosLogger) Level(level string) {
	switch strings.ToLower(level) {
	case "debug":
		logger := l.Logger.Level(zerolog.DebugLevel)
		l.Logger = NewLoggerWith(&logger)
	case "warn":
		logger := l.Logger.Level(zerolog.WarnLevel)
		l.Logger = NewLoggerWith(&logger)
	case "error":
		logger := l.Logger.Level(zerolog.ErrorLevel)
		l.Logger = NewLoggerWith(&logger)
	default:
		logger := l.Logger.Level(zerolog.InfoLevel)
		l.Logger = NewLoggerWith(&logger)
	}
}
