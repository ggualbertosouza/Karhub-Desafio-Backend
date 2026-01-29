package server

import (
	HttpServer "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(env string) *gin.Engine {
	setGinMode(env)

	router := gin.New()
	setMiddlewares(router)
	setRoutes(router)

	return router
}

func setGinMode(env string) {
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}

func setMiddlewares(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
}

func setRoutes(router *gin.Engine) {
	router.GET("/health", healthCheck)
	HttpServer.RegisterRouter(router)
}
