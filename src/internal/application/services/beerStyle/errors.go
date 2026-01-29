package BeerStyleService

import "errors"

var (
	ErrEmptyBeerStyle = errors.New("Empty beer style")
	ErrEmptyList      = errors.New("Beer style list cannot be nil.")
)
