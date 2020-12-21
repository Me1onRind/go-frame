package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

const (
	requestIDKey = "requestID"
)

type tracer struct {
	requestID string
}

func (t *tracer) RequestID() string {
	return t.requestID
}

func setRequestID(c *gin.Context) {
	xRequestID := c.GetHeader("X-Request-ID")
	if len(xRequestID) == 0 {
		xRequestID = uuid.NewV4().String()
	}
	c.Set(requestIDKey, xRequestID)
}

func getTracer(c *gin.Context) *tracer {
	requestID := c.MustGet(requestIDKey).(string)
	return &tracer{
		requestID: requestID,
	}
}
