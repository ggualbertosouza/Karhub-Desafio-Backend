package BeerStyleService

import (
	BeerStyleEntity "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/domain/entities/beerStyle"
	"math"
	"sort"
)

type candidate struct {
	item     BeerStyleList
	distance float64
}

func SelectBeerStyle(
	inputTemp float64,
	list []BeerStyleList,
) (*BeerStyleEntity.BeerStyle, error) {
	if len(list) == 0 {
		return nil, ErrEmptyBeerStyle
	}

	candidates := make([]candidate, 0, len(list))

	for _, style := range list {
		candidates = append(candidates, candidate{
			item:     style,
			distance: math.Abs(style.AvgTemp - inputTemp),
		})
	}

	alphabeticSort(candidates)

	selected := candidates[0].item

	return &BeerStyleEntity.BeerStyle{
		ID:      selected.Id,
		Name:    selected.Name,
		MinTemp: selected.MinTemp,
		MaxTemp: selected.Maxtemp,
	}, nil
}

func alphabeticSort(candidates []candidate) {
	sort.Slice(candidates, func(i, j int) bool {
		if candidates[i].distance == candidates[j].distance {
			return candidates[i].item.Name < candidates[j].item.Name
		}

		return candidates[i].distance < candidates[j].distance
	})
}
