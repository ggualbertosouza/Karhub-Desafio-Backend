package BeerStyleEntity

import (
	"github.com/google/uuid"
)

func New(
	id *uuid.UUID,
	name string,
	minTemp, maxTemp float64,
	tempType *TemperatureType,
) (*BeerStyle, error) {

	validTempType, err := validateBeerStyle(name, minTemp, maxTemp, tempType)
	if err != nil {
		return nil, err
	}

	styleID := uuid.New()
	if id != nil {
		styleID = *id
	}

	return &BeerStyle{
		ID:      styleID,
		Name:    name,
		Type:    validTempType,
		MinTemp: minTemp,
		MaxTemp: maxTemp,
	}, nil
}

func validateBeerStyle(
	name string,
	minTemp, maxTemp float64,
	tempType *TemperatureType,
) (TemperatureType, error) {
	if name == "" {
		return "", ErrNameRequired
	}

	if minTemp >= maxTemp {
		return "", ErrInvalidTemperature
	}

	if tempType == nil {
		return Celsius, nil
	}

	switch *tempType {
	case Celsius, Fahrenheit, Kelvin:
		return *tempType, nil
	default:
		return "", ErrInvalidTemperatureType
	}
}
