package logger

import (
	"fmt"

	"github.com/micro/go-micro/v2/logger"
	"go.uber.org/zap"
)

func SetGoMicroLogger(lg *zap.Logger) {
	logger.DefaultLogger = newGoMicroLogger(lg)
}

type goMicroLogger struct {
	lg *zap.Logger
}

func newGoMicroLogger(lg *zap.Logger) *goMicroLogger {
	return &goMicroLogger{
		lg: lg,
	}
}
func (g *goMicroLogger) Init(options ...logger.Option) error {
	return nil
}

func (g *goMicroLogger) Fields(fields map[string]interface{}) logger.Logger {
	zapFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		zapFields = append(zapFields, zap.Any(k, v))
	}
	return newGoMicroLogger(g.lg.With(zapFields...))
}

func (g *goMicroLogger) Log(level logger.Level, v ...interface{}) {
	switch level {
	case logger.DebugLevel, logger.TraceLevel:
		g.lg.Debug(fmt.Sprintln(v...))
	case logger.InfoLevel:
		g.lg.Info(fmt.Sprintln(v...))
	case logger.WarnLevel:
		g.lg.Warn(fmt.Sprintln(v...))
	case logger.ErrorLevel:
		g.lg.Error(fmt.Sprintln(v...))
	case logger.FatalLevel:
		g.lg.Fatal(fmt.Sprintln(v...))
	default:
		g.lg.Info(fmt.Sprintln(v...))
	}
}

func (g *goMicroLogger) Logf(level logger.Level, format string, v ...interface{}) {
	switch level {
	case logger.DebugLevel, logger.TraceLevel:
		g.lg.Debug(fmt.Sprintf(format, v...))
	case logger.InfoLevel:
		g.lg.Info(fmt.Sprintf(format, v...))
	case logger.WarnLevel:
		g.lg.Warn(fmt.Sprintf(format, v...))
	case logger.ErrorLevel:
		g.lg.Error(fmt.Sprintf(format, v...))
	case logger.FatalLevel:
		g.lg.Fatal(fmt.Sprintf(format, v...))
	default:
		g.lg.Info(fmt.Sprintf(format, v...))
	}
}

func (g *goMicroLogger) Options() logger.Options {
	return logger.Options{}
}

func (g *goMicroLogger) String() string {
	return "zap-go-micro-logger"
}
