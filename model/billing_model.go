package model

import "github.com/google/uuid"

type Billing struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID uuid.UUID
	Amount float64
}
