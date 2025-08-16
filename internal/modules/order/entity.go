package order

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
	Status    string  `json:"status"`
}
