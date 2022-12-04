package model

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
}
