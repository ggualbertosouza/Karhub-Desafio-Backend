package HttpServer

import (
	BeerStyleHandler "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine) {
	rg := router.Group("/beerstyles")

	rg.POST("", BeerStyleHandler.CreateBs)
	rg.PUT("/:id", BeerStyleHandler.UpdateBs)
	rg.DELETE("/:id", BeerStyleHandler.SoftDeletebyId)
	rg.PATCH("/:id/active", BeerStyleHandler.ActiveById)

	rg.GET("/:id", BeerStyleHandler.GetByID)
	rg.GET("", BeerStyleHandler.List)
	rg.GET("/temperature", BeerStyleHandler.GetByTemperature)
}
