package BeerStyleDtos

type BsDefaultOutput struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Mintemp float64 `json:"minTemp"`
	MaxTemp float64 `json:"maxTemp"`
}
