package logger

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"runtime"
	"testing"
	"time"
)

var (
	config *Config
	caller string
	writer *stringWriter
)

type stringWriter struct {
	value string
}

func (s *stringWriter) Write(b []byte) (n int, err error) {
	s.value = string(b)
	return len(b), nil
}

func TestMain(m *testing.M) {
	writer = &stringWriter{}
	config = NewDefaultConfig()
	config.TimeFormat = func(t time.Time) string {
		return "2020-12-04 11:30:00.123"
	}
	config.NewWriter = func(level Level) io.Writer {
		return writer
	}
	m.Run()
}

func getCaller(fixedLine int) string {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		caller = fmt.Sprintf("%s:%d", file, line+fixedLine)
		return caller
	}
	return ""
}

func Test_SimpleOutput(t *testing.T) {
	lg := NewLogger(config)

	lg.WithFields(KV("key", "value")).Debug("hello,world!")
	assert.Empty(t, writer.value)

	lg.WithFields(KV("key", "value")).Info("hello,world!")
	caller = getCaller(-1)
	assert.Equal(t, "[INFO]|2020-12-04 11:30:00.123|"+caller+"|key[value]|hello,world!\n", writer.value)

	lg.WithFields(KV("key", "value")).Error("hello,world!")
	caller = getCaller(-1)
	assert.Equal(t, "[ERROR]|2020-12-04 11:30:00.123|"+caller+"|key[value]|hello,world!\n", writer.value)

	lg.WithFields(KV("key", "value"), KV("k", "v")).Warn("hello,world!")
	caller = getCaller(-1)
	assert.Equal(t, "[WARN]|2020-12-04 11:30:00.123|"+caller+"|key[value]|k[v]|hello,world!\n", writer.value)

	t.Log(writer.value)
}

func Test_WriteToFile(t *testing.T) {
	cfg := NewDefaultConfig()
	writer, err := NewFileRatatelogWriter("./", "test.log")
	if assert.Empty(t, err) {
		cfg.NewWriter = func(level Level) io.Writer {
			return writer
		}
		config.TimeFormat = func(t time.Time) string {
			return "2020-12-04 11:30:00.123"
		}
		lg := NewLogger(cfg)
		lg.Info("Hello, world")
	}
}
