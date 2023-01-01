package repository

import (
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"gorm.io/gorm"
)

type ProductReviewRepo interface {
	Save(m *model.ProductReview) error
}

type productReviewRepo struct {
	db *gorm.DB
}

func NewProductReviewRepo(db *gorm.DB) ProductReviewRepo {
	return &productReviewRepo{db: db}
}
func (p *productReviewRepo) Save(m *model.ProductReview) error {
	return p.db.Save(m).Error
}
