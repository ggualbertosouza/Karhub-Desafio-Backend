package BeerStyleHandler

import (
	"errors"
	BeerStyleEntity "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/domain/beerStyle"
	SpotifyEntity "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/domain/spotify"
	HttpContext "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/http/context"
	InMemoryCache "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/infra/cache/inMemory"
	"math"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type GetByTempRequest struct {
	Temperature float64 `json:"temperature" binding:"required"`
}

func GetByTemperature(ctx *gin.Context) {
	var req GetByTempRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		HttpContext.BadRequest(ctx, err.Error())
		return
	}

	styles, err := InMemoryCache.BsCache.Get(ctx.Request.Context())
	if err != nil {
		HttpContext.BadRequest(ctx, err.Error())
		return
	}

	selected, err := SelectBeerStyleByTemperature(
		req.Temperature,
		styles,
	)
	if err != nil {
		HttpContext.NotFound(ctx, err.Error())
		return
	}

	playlist, _ := getPlaylisyByBsName(selected.Name)

	ctx.JSON(http.StatusOK, gin.H{
		"beerStyle": selected.Name,
		"playlist":  playlist,
	})
}

func SelectBeerStyleByTemperature(
	inputTemp float64,
	list []InMemoryCache.BeerStyleList,
) (*BeerStyleEntity.BeerStyle, error) {
	if len(list) == 0 {
		return nil, errors.New("no beer styles available")
	}

	var (
		selected InMemoryCache.BeerStyleList
		bestDist = math.MaxFloat64
		found    = false
	)

	for _, style := range list {
		dist := math.Abs(style.AvgTemp - inputTemp)

		if !found ||
			dist < bestDist ||
			(dist == bestDist && strings.Compare(style.Name, selected.Name) < 0) {

			selected = style
			bestDist = dist
			found = true
		}
	}

	return BeerStyleEntity.New(
		&selected.Id,
		selected.Name,
		nil,
		selected.MinTemp,
		selected.MaxTemp,
		nil,
	)
}

func getPlaylisyByBsName(name string) (*SpotifyEntity.Playlist, error) {
	playlist, err := InMemoryCache.PlaylistMockCache.Get(name)
	if err != nil {
		return nil, err
	}

	return playlist, nil
}
