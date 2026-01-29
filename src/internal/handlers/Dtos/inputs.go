package BeerStyleDtos

type CreateBsRequest struct {
	Name    string  `json:"name" binding:"required"`
	MinTemp float64 `json:"minTemp" binding:"required"`
	MaxTemp float64 `json:"maxTemp" binding:"required"`
}

type UpdateBsRequest struct {
	Name    *string  `json:"name"`
	MinTemp *float64 `json:"minTemp"`
	MaxTemp *float64 `json:"maxTemp"`
}
