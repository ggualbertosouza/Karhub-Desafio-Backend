package BeerStyleUseCase

import (
	BeerStyleDtos "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/application/Dtos"
	BeerStyleEntity "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/domain/entities/beerStyle"
)

func CreateBS(input *BeerStyleDtos.CreateBSInput) (string, error) {
	bs, err := BeerStyleEntity.New(nil, input.Name, input.MinTemp, input.MaxTemp, nil)
	if err != nil {
		return "", err
	}

	// #TODO Create in database

	return bs.ID.String(), nil
}
