package repository

import "gorm.io/gorm"

type ProductRatingRepo interface {
}

type productRatingRepo struct {
	db *gorm.DB
}

func NewProductRatingRepo(db *gorm.DB) ProductRatingRepo {
	return &productRepo{db: db}
}
