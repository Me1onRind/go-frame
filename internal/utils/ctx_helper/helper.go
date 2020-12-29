package ctx_helper

import (
	"context"
	"github.com/gin-gonic/gin"
	customContext "go-frame/internal/pkg/context"
	//"go-frame/internal/pkg/logger"
)

const (
	commonContextKey = "commonCtx"
	httpContextKey   = "httpCtx"
	customContextKey = "customCtx"
)

func SetCommonContext(c context.Context, ctx *customContext.CommonContext) context.Context {
	return context.WithValue(c, commonContextKey, ctx)
}

func GetCommonContext(c context.Context) *customContext.CommonContext {
	return c.Value(commonContextKey).(*customContext.CommonContext)
}

func SetHttpContext(c *gin.Context, ctx *customContext.HttpContext) {
	c.Set(httpContextKey, ctx)
}

func GetHttpContext(c *gin.Context) *customContext.HttpContext {
	return c.MustGet(httpContextKey).(*customContext.HttpContext)
}

func SetCustomContext(c context.Context, ctx customContext.Context) context.Context {
	return context.WithValue(c, customContextKey, ctx)
}

func GetCustomContext(c context.Context) customContext.Context {
	return c.Value(customContextKey).(customContext.Context)
}
