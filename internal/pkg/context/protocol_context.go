package context

import (
	"context"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go-frame/global"
)

type HttpContext struct {
	Raw interface{}

	*contextS
	*gin.Context
	reqeustID string
}

func NewHttpContext(c *gin.Context) *HttpContext {
	h := &HttpContext{
		contextS: newContextS(),
		Context:  c,
	}
	h.contextS.reqeustID = c.GetString(global.ContextRequestIDKey)
	return h
}

func (h *HttpContext) RequestID() string {
	return h.reqeustID
}

type CommonContext struct {
	*contextS
	context.Context
}

func NewCommonContext(ctx context.Context) *CommonContext {
	c := &CommonContext{
		contextS: newContextS(),
		Context:  ctx,
	}

	val := ctx.Value(global.ContextRequestIDKey)
	if val != nil {
		if requestID, ok := val.(string); ok {
			c.contextS.reqeustID = requestID
		}
	}

	if len(c.contextS.reqeustID) == 0 {
		c.contextS.reqeustID = uuid.NewV4().String()
	}

	return c
}
