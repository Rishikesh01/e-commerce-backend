package repository

import "gorm.io/gorm"

type ProductRepository struct {
	Db *gorm.DB
}

func NewProductRepository(Db *gorm.DB) *ProductRepository {
	return &ProductRepository{Db: Db}
}
