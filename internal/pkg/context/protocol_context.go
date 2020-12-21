package context

import (
	"github.com/gin-gonic/gin"
)

type HttpContext struct {
	Raw interface{}

	*contextS
	*gin.Context
	reqeustID string
}

func NewHttpContext(c *gin.Context) *HttpContext {
	return &HttpContext{
		contextS: newContextS(),
		Context:  c,
	}
}

func (h *HttpContext) RequestID() string {
	return h.reqeustID
}
