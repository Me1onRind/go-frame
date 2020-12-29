package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go-frame/internal/pkg/logger"
	"go-frame/internal/utils/ctx_helper"
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
			logger.Errorf("Get request body error:%v", err)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(request))

		lw := &logWriter{
			ResponseWriter: c.Writer,
			buff:           &bytes.Buffer{},
		}
		c.Writer = lw

		logger.WithTrace(ctx).WithFields(
			logger.KV("method", c.Request.Method),
			logger.KV("URI", c.Request.URL.Path),
			logger.KV("rawQuery", c.Request.URL.RawQuery),
			logger.KV("request", string(request)),
			logger.KV("clientIP", c.ClientIP()),
		).Info("HTTP Request Start")

		start := time.Now()
		c.Next()
		end := time.Now()

		logger.WithTrace(ctx).WithFields(
			logger.KV("method", c.Request.Method),
			logger.KV("URI", c.Request.URL.Path),
			logger.KV("rawQuery", c.Request.URL.RawQuery),
			logger.KV("request", string(request)),
			logger.KV("clientIP", c.ClientIP()),
			logger.KV("response", string(lw.buff.Bytes())),
			logger.KV("cost", end.Sub(start)),
		).Info("HTTP Request End")
	}
}
