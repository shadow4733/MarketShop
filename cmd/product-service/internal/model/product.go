package model

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ProductId       uuid.UUID `json:"product_id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Price           float64   `json:"price"`
	CategoryIDs     []string  `json:"category_ids"`
	StockQuantity   int32     `json:"stock_quantity"`
	Rating          float64   `json:"rating"`
	DiscountPercent float64   `json:"discount_percent"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
