package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go-frame/global"
	"go-frame/internal/pkg/logger"
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

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		xRequestID := getRequestID(c)
		c.Set(global.ContextRequestIDKey, xRequestID)

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

		logger.WithTrace(getTracer(c)).WithFields(
			logger.KV("method", c.Request.Method),
			logger.KV("URI", c.Request.URL.Path),
			logger.KV("rawQuery", c.Request.URL.RawQuery),
			logger.KV("request", string(request)),
			logger.KV("clientIP", c.ClientIP()),
		).Info("Request begin")

		start := time.Now()
		c.Next()
		end := time.Now()

		logger.WithTrace(getTracer(c)).WithFields(
			logger.KV("method", c.Request.Method),
			logger.KV("URI", c.Request.URL.Path),
			logger.KV("rawQuery", c.Request.URL.RawQuery),
			logger.KV("request", string(request)),
			logger.KV("clientIP", c.ClientIP()),
			logger.KV("response", string(lw.buff.Bytes())),
			logger.KV("cost", end.Sub(start)),
		).Info("Request completed")
	}
}
