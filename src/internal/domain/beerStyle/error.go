package BeerStyleEntity

import "errors"

var (
	ErrNameRequired           = errors.New("Beer Style name is required.")
	ErrInvalidTemperature     = errors.New("Invalid temperature range.")
	ErrInvalidTemperatureType = errors.New("Invalid temperature type.")
)
