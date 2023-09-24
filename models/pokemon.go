package models

type Pokemon struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	Weight int    `json:"weight"`
	Height int    `json:"height"`
}
