package klog

type KcronLogger struct {
	Logger *Logger
}

func (w *KcronLogger) Info(msg string, keysAndValues ...interface{}) {
	w.Logger.Info().Interface("details", keysAndValues).Msg(msg)
}

func (w *KcronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	w.Logger.Error().Interface("details", keysAndValues).Err(err).Msg(msg)
}

func (w *KcronLogger) Log(keyvals ...interface{}) error {
	w.Logger.Info().Interface("details", keyvals).Send()
	return nil
}
