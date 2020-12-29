package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go-frame/internal/pkg/logger"
	"go-frame/internal/utils/ctx_helper"
	"go.uber.org/zap"
	"io/ioutil"
	"time"
)

type logWriter struct {
	gin.ResponseWriter
	buff *bytes.Buffer
}

func (w *logWriter) Write(b []byte) (int, error) {
	w.buff.Write(b)
	return w.ResponseWriter.Write(b)
}

func AccessLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := ctx_helper.GetHttpContext(c)
		request, err := c.GetRawData()
		if err != nil {
			ctx.Logger().Error("Get request body error", zap.Error(err))
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(request))

		lw := &logWriter{
			ResponseWriter: c.Writer,
			buff:           &bytes.Buffer{},
		}
		c.Writer = lw

		ctx.Logger().Info("HTTP Request Start",
			zap.String("method", c.Request.Method),
			zap.String("PATH", c.Request.URL.Path),
			zap.String("rawQuery", c.Request.URL.RawQuery),
			zap.String("reqBody", string(request)),
			zap.String("clientIP", c.ClientIP()),
		)

		start := time.Now()
		c.Next()
		end := time.Now()

		ctx.Logger().Info("HTTP Request End",
			zap.String("method", c.Request.Method),
			zap.String("PATH", c.Request.URL.Path),
			zap.String("rawQuery", c.Request.URL.RawQuery),
			zap.String("reqBody", string(request)),
			zap.String("clientIP", c.ClientIP()),
			zap.String("resp", string(lw.buff.Bytes())),
			zap.Duration("cost", end.Sub(start)),
		)
	}
}
