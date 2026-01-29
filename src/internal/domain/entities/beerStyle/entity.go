package BeerStyleEntity

import (
	"github.com/google/uuid"
)

type TemperatureType string

const (
	Celsius    TemperatureType = "Celsius"
	Fahrenheit TemperatureType = "Fahrenheit"
	Kelvin     TemperatureType = "Kelvin"
)

type BeerStyle struct {
	ID      uuid.UUID
	Name    string
	Type    TemperatureType
	MinTemp float64
	MaxTemp float64
}

func (b *BeerStyle) AverageTemperature() float64 {
	return (b.MinTemp + b.MaxTemp) / 2
}
