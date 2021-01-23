package middleware

import (
	"bytes"
	"go-frame/internal/core/context"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type logWriter struct {
	gin.ResponseWriter
	buff *bytes.Buffer
}

func (w *logWriter) Write(b []byte) (int, error) {
	w.buff.Write(b)
	return w.ResponseWriter.Write(b)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.GetFromGinContext(c)
		request, err := c.GetRawData()
		if err != nil {
			ctx.Logger.Error("Get request body error", zap.Error(err))
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(request))

		lw := &logWriter{
			ResponseWriter: c.Writer,
			buff:           &bytes.Buffer{},
		}
		c.Writer = lw

		start := time.Now()
		defer func() {
			end := time.Now()
			ctx.Logger.Info("access log",
				zap.String("method", c.Request.Method),
				zap.String("path", c.Request.URL.Path),
				zap.String("rawQuery", c.Request.URL.RawQuery),
				zap.String("reqBody", string(request)),
				zap.String("clientIP", c.ClientIP()),
				zap.String("resp", lw.buff.String()),
				zap.Duration("cost", end.Sub(start)),
			)
		}()

		c.Next()

	}
}
