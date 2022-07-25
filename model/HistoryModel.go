package model

import "github.com/google/uuid"

type UserOrderHistory struct {
	ItemID    uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Quantity  uint16
	ProductID uuid.UUID
	Price     float64
}
