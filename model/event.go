package model

type Event struct {
	Price int    `json:"price"`
	Stock int    `json:"stock"`
	Id    int    `json:"id"`
	Title string `json:"title"`
}
