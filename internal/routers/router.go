package routers

import (
	"github.com/gin-gonic/gin"
	"go-frame/internal/pkg/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	apiRouter := r.Group("/api")
	registerUserApi(apiRouter)
	registerJWTApi(apiRouter)

	return r
}
