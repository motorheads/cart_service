package models

type Cart struct {
	UserID       int `json:"user_id"`
	ProductID    int `json:"product_id"`
	ProductCount int `json:"product_count"`
}
