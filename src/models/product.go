package models

type Product struct {
	Id       uint   `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Price    int64  `json:"price,omitempty"`
	Category string `json:"category,omitempty"`
}