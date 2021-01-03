package routers

import (
	"github.com/gin-gonic/gin"
	"go-frame/internal/core/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Tracing())
	r.Use(middleware.AccessLogger())
	r.Use(middleware.Recover())

	apiRouter := r.Group("/api")
	registerUserApi(apiRouter)
	registerJWTApi(apiRouter)

	return r
}
