package model

import "github.com/google/uuid"

type Product struct {
	Id            uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name          string
	ImageLocation string
	Price         float64
}
