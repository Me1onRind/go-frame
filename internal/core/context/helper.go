package context

import "context"

type key uint8

const (
	customContextKey key = iota
)

func LoadIntoContext(ctx *Context, libCtx context.Context) context.Context {
	return context.WithValue(libCtx, customContextKey, ctx)
}

func GetFromContext(libCtx context.Context) *Context {
	return libCtx.Value(customContextKey).(*Context)
}
