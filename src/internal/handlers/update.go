package BeerStyleHandler

import (
	BeerStyleDtos "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/handlers/Dtos"
	HttpContext "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/http/context"
	InMemoryCache "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/infra/cache/inMemory"
	BsRepository "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/infra/repositories"
	"github/ggualbertosouza/Karhub-Desafio-Backend/src/pkg/postgres"

	"github.com/gin-gonic/gin"
)

func UpdateBs(ctx *gin.Context) {
	var req BeerStyleDtos.UpdateBsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		HttpContext.BadRequest(ctx, "Invalid Payload")
		return
	}

	bsId := ctx.Param("id")
	if bsId == "" {
		HttpContext.BadRequest(ctx, "Invalid Id")
		return
	}

	if req.Name == nil &&
		req.MinTemp == nil &&
		req.MaxTemp == nil {
		HttpContext.BadRequest(ctx, "Invalid Payload.")
		return
	}

	queryRepo := BsRepository.NewQuery(postgres.PostgresDb)
	cmdRepo := BsRepository.NewCmd(postgres.PostgresDb)

	_, err := queryRepo.GetById(ctx, bsId)
	if err != nil {
		HttpContext.NotFound(ctx, err.Error())
		return
	}

	err = cmdRepo.Update(ctx, bsId, req.Name, req.MinTemp, req.MaxTemp)
	if err != nil {
		HttpContext.BadRequest(ctx, err.Error())
		return
	}

	InMemoryCache.BsCache.Populate(ctx)
	HttpContext.ResourceUpdated(ctx, "Beer Style")
}
