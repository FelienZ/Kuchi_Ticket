package model

type Create struct {
	Price int    `json:"price"`
	Stock int    `json:"stock"`
	Title string `json:"title"`
}
