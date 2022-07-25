package model

import "github.com/google/uuid"

type ProductSeller struct {
	SellerId  uuid.UUID `gorm:"primaryKey"`
	ProductId uuid.UUID `gorm:"primaryKey"`
}
