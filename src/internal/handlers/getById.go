package BeerStyleHandler

import (
	BeerStyleDtos "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/handlers/Dtos"
	HttpContext "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/http/context"
	BsRepository "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/infra/repositories"
	"github/ggualbertosouza/Karhub-Desafio-Backend/src/pkg/postgres"

	"github.com/gin-gonic/gin"
)

func GetByID(ctx *gin.Context) {
	bsId := ctx.Param("id")
	if bsId == "" {
		HttpContext.BadRequest(ctx, "Id is required.")
	}

	queryRepo := BsRepository.NewQuery(postgres.PostgresDb)
	beerStyle, err := queryRepo.GetById(ctx, bsId)
	if err != nil {
		HttpContext.NotFound(ctx, err.Error())
		return
	}

	if beerStyle.Active == false {
		HttpContext.BadRequest(ctx, "Unavailable Beer Style.")
		return
	}

	HttpContext.ResponseOk(ctx, mapToOutput(beerStyle))
}

func mapToOutput(fromDb *BsRepository.BsModel) *BeerStyleDtos.BsDefaultOutput {
	return &BeerStyleDtos.BsDefaultOutput{
		Id:      fromDb.Id.String(),
		Name:    fromDb.Name,
		Mintemp: fromDb.MinTemp,
		MaxTemp: fromDb.MaxTemp,
	}
}
