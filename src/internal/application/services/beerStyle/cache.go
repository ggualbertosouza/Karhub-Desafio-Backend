package BeerStyleService

import BeerStyleEntity "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/domain/entities/beerStyle"

func MapToCache(list *[]BeerStyleEntity.BeerStyle) ([]BeerStyleList, error) {
	if list == nil {
		return nil, ErrEmptyList
	}

	cache := make([]BeerStyleList, 0, len(*list))

	for _, style := range *list {
		cache = append(cache, BeerStyleList{
			Id:      style.ID,
			Name:    style.Name,
			MinTemp: style.MinTemp,
			Maxtemp: style.MaxTemp,
			AvgTemp: style.AverageTemperature(),
		})
	}

	return cache, nil
}
