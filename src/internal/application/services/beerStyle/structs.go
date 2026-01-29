package BeerStyleService

import "github.com/google/uuid"

type BeerStyleList struct {
	Id      uuid.UUID
	Name    string
	MinTemp float64
	Maxtemp float64
	AvgTemp float64
}
