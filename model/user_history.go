package model

import "github.com/google/uuid"

type UserOrderHistory struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	BillingID uint64
	ProductID uuid.UUID
}
