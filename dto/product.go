package dto

import (
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"github.com/google/uuid"
)

type ProductSearch struct {
	ID          uuid.UUID             `json:"id"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Sellers     []model.ProductSeller `json:"sellers"`
}

type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       uint64    `json:"price"`
}
