package BeerStyleHandler

import (
	HttpContext "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/http/context"
	BsRepository "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/infra/repositories"
	"github/ggualbertosouza/Karhub-Desafio-Backend/src/pkg/postgres"

	"github.com/gin-gonic/gin"
)

func List(ctx *gin.Context) {
	queryRepo := BsRepository.NewQuery(postgres.PostgresDb)
	bsList, err := queryRepo.ListAll(ctx)
	if err != nil {
		HttpContext.BadRequest(ctx, err.Error())
		return
	}

	HttpContext.ResponseOk(ctx, bsList)
}
