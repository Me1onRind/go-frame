package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"runtime"
	"time"
)

type Level uint8

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	default:
		return "ERROR"
	}
}

type Field struct {
	Key   string
	Value interface{}
}

func KV(key string, value interface{}) *Field {
	return &Field{key, value}
}

func JSONKV(key string, value interface{}) *Field {
	jsonStr, _ := json.Marshal(value)
	return &Field{key, string(jsonStr)}
}

type Logger struct {
	fields     []*Field
	cfg        *Config
	callerSkip int
}

func NewLogger(cfg *Config) *Logger {
	return &Logger{
		cfg:        cfg,
		callerSkip: 3,
	}
}

func (l *Logger) clone() *Logger {
	nl := *l
	if l.fields != nil {
		nl.fields = make([]*Field, len(l.fields), len(l.fields)+4)
		copy(nl.fields, l.fields)
	}
	return &nl
}

func (l *Logger) Close() {
}

func (l *Logger) WithFields(fs ...*Field) *Logger {
	ll := l.clone()
	ll.fields = append(ll.fields, fs...)
	return ll
}

func (l *Logger) WithTrace(tracer Tracer) *Logger {
	return l.WithFields(KV("requestID", tracer.RequestID()))
}

func (l *Logger) output(level Level, message string) {
	cfg := l.cfg
	separate := cfg.Separate

	if level < cfg.Level {
		return
	}

	writer := l.cfg.NewWriter(level)
	builder := new(bytes.Buffer)
	// log level
	writeSingleString(builder, level.String())
	// log time
	now := time.Now()
	builder.WriteRune(l.cfg.Separate)
	builder.WriteString(l.cfg.TimeFormat(now))
	builder.WriteRune(separate)
	// log caller
	_, file, line, ok := runtime.Caller(l.callerSkip)
	if ok {
		builder.WriteString(fmt.Sprintf("%s:%d", file, line))
	}
	builder.WriteRune(separate)

	for _, field := range l.fields {
		builder.WriteString(fmt.Sprintf("%s[%v]", field.Key, field.Value))
		builder.WriteRune(separate)
	}
	builder.WriteString(message)
	builder.WriteByte('\n')
	writer.Write(builder.Bytes())
}

func (l *Logger) SetCallerSkip(skip int) {
	l.callerSkip = skip
}

func (l *Logger) Debug(v ...interface{}) {
	l.output(LevelDebug, fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.output(LevelDebug, fmt.Sprintf(format, v...))
}

func (l *Logger) Info(v ...interface{}) {
	l.output(LevelInfo, fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.output(LevelInfo, fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(v ...interface{}) {
	l.output(LevelWarn, fmt.Sprint(v...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.output(LevelWarn, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.output(LevelError, fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.output(LevelError, fmt.Sprintf(format, v...))
}
