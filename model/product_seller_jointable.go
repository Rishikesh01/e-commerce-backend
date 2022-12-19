package model

import "github.com/google/uuid"

type ProductSeller struct {
	ID        uint      `json:"id"`
	ProductID uuid.UUID `json:"product_id"`
	SellerID  uuid.UUID `json:"seller_id"`
	Price     uint64    `json:"price"`
}
