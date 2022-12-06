package model

import "github.com/google/uuid"

type ProductSeller struct {
	ProductID uuid.UUID `gorm:"primaryKey"`
	SellerID  uuid.UUID `gorm:"primaryKey"`
	Price     uint64
}
