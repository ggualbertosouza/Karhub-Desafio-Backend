package server

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(env string) *gin.Engine {
	setGinMode(env)

	router := gin.New()

	router.GET("/health", healthCheck)

	return router
}

func setGinMode(env string) {
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}
