package model

import "github.com/google/uuid"

type ProductSeller struct {
	ProductID uuid.UUID `gorm:"primaryKey" json:"product_id"`
	SellerID  uuid.UUID `gorm:"primaryKey" json:"seller_id"`
	Price     uint64    `json:"price"`
}
