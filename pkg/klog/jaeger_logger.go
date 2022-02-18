package klog

type JaegerLogger struct {
	Logger *Logger
}

// Error logs a message at error priority
func (l *JaegerLogger) Error(msg string) {
	l.Logger.Error().Msg(msg)
}

// Infof logs a message at info priority
func (l *JaegerLogger) Infof(msg string, args ...interface{}) {
	l.Logger.Info().Interface("details", args).Msg(msg)
}

// Infof logs a message at info priority
func (l *JaegerLogger) Debugf(msg string, args ...interface{}) {
	l.Logger.Debug().Interface("details", args).Msg(msg)
}

func (l *JaegerLogger) Log(keyvals ...interface{}) error {
	l.Logger.Info().Interface("details", keyvals).Send()
	return nil
}
