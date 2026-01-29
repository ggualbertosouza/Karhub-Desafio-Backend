package BeerStyleHandler

import (
	BeerStyleEntity "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/domain/beerStyle"
	BeerStyleDtos "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/handlers/Dtos"
	HttpContext "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/http/context"
	BsRepository "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/infra/repositories"
	"github/ggualbertosouza/Karhub-Desafio-Backend/src/pkg/postgres"

	"github.com/gin-gonic/gin"
)

func CreateBs(ctx *gin.Context) {
	var req BeerStyleDtos.CreateBsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		HttpContext.BadRequest(ctx, err.Error())
		return
	}

	bs, err := BeerStyleEntity.New(nil, req.Name, nil, req.MinTemp, req.MaxTemp, nil)
	if err != nil {
		HttpContext.BadRequest(ctx, err.Error())
		return
	}

	cmdRepo := BsRepository.NewCmd(postgres.PostgresDb)
	err = cmdRepo.Create(ctx, bs)
	if err != nil {
		HttpContext.BadRequest(ctx, err.Error())
		return
	}

	HttpContext.ResourceCreated(ctx, "Beer Style")
}
