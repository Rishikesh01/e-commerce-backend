package model

import "github.com/google/uuid"

type Billing struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID uuid.UUID `json:"user_id"`
	Amount uint64    `json:"amount"`
}
