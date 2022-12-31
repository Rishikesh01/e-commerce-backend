package repository

import "gorm.io/gorm"

type ProductReviewRepo interface {
}

type productReviewRepo struct {
	db *gorm.DB
}

func NewProductReviewRepo(db *gorm.DB) ProductReviewRepo {
	return &productReviewRepo{db: db}
}
