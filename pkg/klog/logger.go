package klog

import (
	"io"
	"os"
	"strings"

	"github.com/guojiangli/picasso/pkg/config"
	"github.com/rs/zerolog"
)

// Logger is the only Logger in log
var defaultLogger = NewLogger()

type Logger struct {
	*zerolog.Logger
	Writer io.Writer
}

func (l *Logger) Log(keyvals ...interface{}) error {
	l.WithLevel(l.GetLevel()).Interface("details", keyvals).Send()
	return nil
}

func (l *Logger) SetLevel(level Level) *Logger {
	l1 := l.Logger.Level(zerolog.Level(level))
	return NewLoggerWith(&l1)
}

func NewLoggerWith(zlogger *zerolog.Logger) *Logger {
	return &Logger{zlogger, defaultLogger.Writer}
}

func (l *Logger) NewLogger() *Logger {
	logger := l.With().Logger()
	return &Logger{&logger, l.Writer}
}

// AutoLevel ...
func (l *Logger) AutoLevel(confKey string) {
	config.OnChange(func(config *config.Config) {
		lvText := strings.ToLower(config.GetString(confKey + ".level"))
		if lvText != "" {
			lv, err := LevelUnmarshalText(lvText)
			if err != nil {
				l.Error().Err(err).Send()
			}
			if zerolog.Level(lv) != l.GetLevel() {
				newl := l.Level(zerolog.Level(lv))
				l.Logger = &newl
				l.Info().Str("level", lvText).Msg("update level")
			}
		}
	})
}

// InitLogger is to initialize Logger
func InitLogger(opts ...*Option) {
	defaultLogger = NewLogger(opts...)
}

func NewLogger(opts ...*Option) *Logger {
	LoggerOpt := defaultOption().MergeOption(opts...)
	var zLogger zerolog.Logger
	if LoggerOpt.Writer == nil {
		LoggerOpt.Writer = os.Stderr
	}
	zLogger = zerolog.New(LoggerOpt.Writer).With().Timestamp().Logger()
	if LoggerOpt.Level == 0 {
		zLogger = zLogger.Level(zerolog.InfoLevel)
	} else {
		zLogger = zLogger.Level(zerolog.Level(LoggerOpt.Level))
	}
	return &Logger{&zLogger, LoggerOpt.Writer}
}
