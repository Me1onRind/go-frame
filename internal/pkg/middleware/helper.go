package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go-frame/global"
	"go-frame/internal/pkg/logger"
)

func setRequestID(c *gin.Context) {
	xRequestID := c.GetHeader(global.ProtocolRequestIDKey)
	if len(xRequestID) == 0 {
		xRequestID = uuid.NewV4().String()
	}
	c.Set(global.ContextRequestIDKey, xRequestID)
}

func getTracer(c *gin.Context) *logger.SimpleTracer {
	requestID := c.MustGet(global.ContextRequestIDKey).(string)
	return &logger.SimpleTracer{
		ReqID: requestID,
	}
}
