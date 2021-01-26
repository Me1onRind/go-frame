package custom_ctx

import (
	"context"

	"github.com/gin-gonic/gin"
)

type key uint8

const (
	ginCustomContextKey     = "gcck"
	customContextKey    key = iota
)

func LoadIntoContext(ctx *Context, libCtx context.Context) context.Context {
	return context.WithValue(libCtx, customContextKey, ctx)
}

func GetFromContext(libCtx context.Context) *Context {
	return libCtx.Value(customContextKey).(*Context)
}

func LoadIntoGinContext(ctx *Context, c *gin.Context) {
	c.Set(ginCustomContextKey, ctx)
}

func GetFromGinContext(c *gin.Context) *Context {
	v, _ := c.Get(ginCustomContextKey)
	return v.(*Context)
}
