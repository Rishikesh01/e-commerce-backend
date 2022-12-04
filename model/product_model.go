package model

import uuid "github.com/jackc/pgtype/ext/gofrs-uuid"

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name        string
	Description string
	SellerID    []Seller
}
