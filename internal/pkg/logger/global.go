package logger

var logger *Logger

func SetLogger(l *Logger) {
	logger = l
}

func WithTrace(t Tracer) *Logger {
	nl := logger.WithTrace(t)
	nl.callerSkip--
	return nl
}

func WithFields(fields ...*Field) *Logger {
	nl := logger.WithFields(fields...)
	nl.callerSkip--
	return nl
}

func Debug(v ...interface{}) {
	logger.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	logger.Debugf(format, v...)
}

func Info(v ...interface{}) {
	logger.Info(v...)
}

func Infof(format string, v ...interface{}) {
	logger.Infof(format, v...)
}

func Warn(v ...interface{}) {
	logger.Warn(v...)
}

func Warnf(format string, v ...interface{}) {
	logger.Warnf(format, v...)
}

func Error(v ...interface{}) {
	logger.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	logger.Errorf(format, v...)
}
