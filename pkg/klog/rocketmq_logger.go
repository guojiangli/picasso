package klog

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
)

type RocketmqLogger struct {
	Logger *Logger
}

func (w *RocketmqLogger) Log(keyvals ...interface{}) error {
	w.Logger.Info().Interface("details", keyvals)
	return nil
}

func (w *RocketmqLogger) Debug(msg string, fields map[string]interface{}) {
	w.Logger.Debug().Interface("details", fields).Msg(msg)
}

func (w *RocketmqLogger) Info(msg string, fields map[string]interface{}) {
	w.Logger.Info().Interface("details", fields).Msg(msg)
}

func (w *RocketmqLogger) Warning(msg string, fields map[string]interface{}) {
	w.Logger.Warn().Interface("details", fields).Msg(msg)
}

func (w *RocketmqLogger) Error(msg string, fields map[string]interface{}) {
	w.Logger.Error().Interface("details", fields).Msg(msg)
}

func (w *RocketmqLogger) Fatal(msg string, fields map[string]interface{}) {
	w.Logger.Fatal().Interface("details", fields).Msg(msg)
}

func (w *RocketmqLogger) Level(level string) {
	switch strings.ToLower(level) {
	case "debug":
		logger := w.Logger.Level(zerolog.DebugLevel)
		w.Logger = NewLoggerWith(&logger)
	case "warn":
		logger := w.Logger.Level(zerolog.WarnLevel)
		w.Logger = NewLoggerWith(&logger)
	case "error":
		logger := w.Logger.Level(zerolog.ErrorLevel)
		w.Logger = NewLoggerWith(&logger)
	default:
		logger := w.Logger.Level(zerolog.InfoLevel)
		w.Logger = NewLoggerWith(&logger)
	}
}

func (w *RocketmqLogger) OutputPath(path string) (err error) {
	var file *os.File
	file, err = os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	logger := w.Logger.Output(file)
	w.Logger = NewLoggerWith(&logger)
	return
}
