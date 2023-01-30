package dto

import (
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"github.com/google/uuid"
)

type ProductSearch struct {
	ID          uuid.UUID             `json:"id"`
	Img         string                `json:"image"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Sellers     []model.ProductSeller `json:"sellers"`
}

type Product struct {
	ID          uuid.UUID `json:"id"`
	Img         string    `json:"image"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Rating      float32   `json:"rating"`
	Price       uint64    `json:"price"`
}

type ProductRatingByUser struct {
	ID     uuid.UUID `json:"id"`
	Rating int       `json:"rating"`
	UserID uuid.UUID `json:"-"`
}

type ProductReview struct {
	ID     uuid.UUID `json:"id"`
	Review string    `json:"review"`
	UserID uuid.UUID `json:"-"`
}
