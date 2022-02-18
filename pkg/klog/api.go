package klog

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

func DefaultLogger() *Logger {
	return defaultLogger
}

func Info(args ...interface{}) {
	DefaultLogger().Info().Msg(fmt.Sprint(args...))
}

func Warn(args ...interface{}) {
	DefaultLogger().Warn().Msg(fmt.Sprint(args...))
}

func Debug(args ...interface{}) {
	DefaultLogger().Debug().Msg(fmt.Sprint(args...))
}

func Error(args ...interface{}) {
	DefaultLogger().Error().Msg(fmt.Sprint(args...))
}

//Fatal方法调用后直接os.Exit(1),慎用
func Fatal(args ...interface{}) {
	DefaultLogger().Fatal().Msg(fmt.Sprint(args...))
}

//Panic方法调用后直接Panic,慎用
func Panic(args ...interface{}) {
	DefaultLogger().Panic().Msg(fmt.Sprint(args...))
}

func Infof(template string, args ...interface{}) {
	DefaultLogger().Info().Msgf(template, args...)
}

func Warnf(template string, args ...interface{}) {
	DefaultLogger().Warn().Msgf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	DefaultLogger().Error().Msgf(template, args...)
}

func Debugf(template string, args ...interface{}) {
	DefaultLogger().Debug().Msgf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	DefaultLogger().Fatal().Msgf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	DefaultLogger().Panic().Msgf(template, args...)
}

func Infos(args ...interface{}) {
	var str string
	for _, s := range args {
		str += spew.Sprintf("［%#+v］", s)
	}
	DefaultLogger().Info().Msg(str)
}
func Warns(args ...interface{}) {
	var str string
	for _, s := range args {
		str += spew.Sprintf("［%#+v］", s)
	}
	DefaultLogger().Warn().Msg(str)
}

func Debugs(args ...interface{}) {
	var str string
	for _, s := range args {
		str += spew.Sprintf("［%#+v］", s)
	}
	DefaultLogger().Debug().Msg(str)
}

func Errors(args ...interface{}) {
	var str string
	for _, s := range args {
		str += spew.Sprintf("［%#+v］", s)
	}
	DefaultLogger().Error().Msg(str)
}

func Fatals(args ...interface{}) {
	var str string
	for _, s := range args {
		str += spew.Sprintf("［%#+v］", s)
	}
	DefaultLogger().Fatal().Msg(str)
}

func Panics(args ...interface{}) {
	var str string
	for _, s := range args {
		str += spew.Sprintf("［%#+v］", s)
	}
	DefaultLogger().Panic().Msg(str)
}
