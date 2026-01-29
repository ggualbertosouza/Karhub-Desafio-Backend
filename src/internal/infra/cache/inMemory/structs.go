package InMemoryCache

import "github.com/google/uuid"

type BeerStyleList struct {
	Id      uuid.UUID
	Name    string
	MinTemp float64
	MaxTemp float64
	AvgTemp float64
}
