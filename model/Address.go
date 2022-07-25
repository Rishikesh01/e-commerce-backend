package model

import "github.com/google/uuid"

type Address struct {
	Id          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	AddressLine string
	LandMark    string
	City        string
	State       string
	ZipCode     int64
	Country     string
	SellerId    uuid.UUID `gorm:"type:uuid"`
	UserId      uuid.UUID `gorm:"type:uuid"`
}
