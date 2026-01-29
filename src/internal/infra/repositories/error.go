package BsRepository

import "errors"

var (
	ErrBeerStyleAlreadyExists = errors.New("beer style already exists")
	ErrBeerStyleNotFound      = errors.New("beer style not found")
)
