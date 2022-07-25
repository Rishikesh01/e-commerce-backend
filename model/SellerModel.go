package model

import "github.com/google/uuid"

type Seller struct {
	Id       uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name     string
	Email    string
	Password string
	Address  Address   `gorm:"foreignKey:SellerId"`
	Product  []Product `gorm:"many2many:product_seller;"`
}
