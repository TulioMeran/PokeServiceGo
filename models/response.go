package models

type Response struct {
	Count   int       `json:"count"`
	Results []Pokemon `json:"results"`
}
