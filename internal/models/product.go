package models

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Photo       string `json:"photo_url"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}
