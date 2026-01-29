package BsRepository

import "github.com/google/uuid"

type BsModel struct {
	Id       uuid.UUID
	Name     string
	Active   bool
	MinTemp  float64
	MaxTemp  float64
	TempType string
}
