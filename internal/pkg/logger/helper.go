package logger

import (
	"bytes"
)

func writeSingleString(builder *bytes.Buffer, value string) {
	builder.WriteRune('[')
	builder.WriteString(value)
	builder.WriteRune(']')
}
