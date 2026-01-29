package BeerStyleDtos

type CreateBSInput struct {
	Name    string  `json:"name"`
	MinTemp float64 `json:"minTemp"`
	MaxTemp float64 `json:"maxTemp"`
}
