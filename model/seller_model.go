package model

import "github.com/google/uuid"

type Seller struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name          string
	Email         string
	BusinessName  string
	Password      string
	ProductSeller []ProductSeller
}
