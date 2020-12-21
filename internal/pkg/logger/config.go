package logger

import (
	"io"
	"os"
	"time"
)

type WriterFunc func(Level) io.Writer
type TimeFormatFunc func(time.Time) string

func DefaultWriter(level Level) io.Writer {
	return os.Stdout
}

func DefaultTimeFormat(now time.Time) string {
	return now.Format("2006-01-02 15:04:05.000")
}

type Config struct {
	NewWriter  WriterFunc
	TimeFormat TimeFormatFunc
	Level      Level
	Separate   rune
}

func NewDefaultConfig() *Config {
	return &Config{
		NewWriter:  DefaultWriter,
		TimeFormat: DefaultTimeFormat,
		Level:      LevelInfo,
		Separate:   '|',
	}
}
