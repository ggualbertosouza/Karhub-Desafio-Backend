package InMemoryCache

import (
	BeerStyleEntity "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/domain/beerStyle"
	BsRepository "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/infra/repositories"
)

func mapToEntity(model BsRepository.BsModel) (*BeerStyleEntity.BeerStyle, error) {
	tempType := BeerStyleEntity.TemperatureType(model.TempType)
	active := model.Active

	return BeerStyleEntity.New(
		&model.Id,
		model.Name,
		&active,
		model.MinTemp,
		model.MaxTemp,
		&tempType,
	)
}

func mapEntityToCache(bs *BeerStyleEntity.BeerStyle) BeerStyleList {
	return BeerStyleList{
		Id:      bs.ID,
		Name:    bs.Name,
		MinTemp: bs.MinTemp,
		MaxTemp: bs.MaxTemp,
		AvgTemp: bs.AverageTemperature(),
	}
}

func mapToCache(list []BsRepository.BsModel) ([]BeerStyleList, error) {
	cache := make([]BeerStyleList, 0, len(list))

	for _, model := range list {
		entity, err := mapToEntity(model)
		if err != nil {
			return nil, err
		}

		cache = append(cache, mapEntityToCache(entity))
	}

	return cache, nil
}
