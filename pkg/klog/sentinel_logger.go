package klog

type SentinelLogger struct {
	Logger *Logger
}

func (w *SentinelLogger) Info(msg string, keysAndValues ...interface{}) {
	w.Logger.Info().Interface("details", keysAndValues).Msg(msg)
}

func (w *SentinelLogger) InfoEnabled() bool {
	return true
}

func (w *SentinelLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	w.Logger.Error().Interface("details", keysAndValues).Err(err).Msg(msg)
}

func (w *SentinelLogger) ErrorEnabled() bool {
	return true
}

func (w *SentinelLogger) Warn(msg string, keysAndValues ...interface{}) {
	w.Logger.Warn().Interface("details", keysAndValues).Msg(msg)
}

func (w *SentinelLogger) WarnEnabled() bool {
	return true
}

func (w *SentinelLogger) Debug(msg string, keysAndValues ...interface{}) {
	w.Logger.Debug().Interface("details", keysAndValues).Msg(msg)
}

func (w *SentinelLogger) DebugEnabled() bool {
	return true
}

func (w *SentinelLogger) Log(keyvals ...interface{}) error {
	w.Logger.Info().Interface("details", keyvals).Send()
	return nil
}
