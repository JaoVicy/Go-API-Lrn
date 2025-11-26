package models

type pizza struct {
	ID    string  `json:"id"`
	Title string  `json:"title"`
	Size  string  `json:"size"`
	Price float64 `json:"price"`
}
