package BeerStyleHandler

import (
	HttpContext "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/http/context"
	BsRepository "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/infra/repositories"
	"github/ggualbertosouza/Karhub-Desafio-Backend/src/pkg/postgres"

	"github.com/gin-gonic/gin"
)

func ActiveById(ctx *gin.Context) {
	bsId := ctx.Param("id")
	if bsId == "" {
		HttpContext.BadRequest(ctx, "Id is required.")
	}
	cmdRepo := BsRepository.NewCmd(postgres.PostgresDb)

	err := cmdRepo.SetActive(ctx, bsId, true)
	if err != nil {
		HttpContext.NotFound(ctx, err.Error())
		return
	}

	HttpContext.ResourceUpdated(ctx, "Beer Style")
}
