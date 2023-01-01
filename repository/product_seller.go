package repository

import (
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"gorm.io/gorm"
)

type ProductSellerRepo interface {
	Save(*model.ProductSeller) error
}

type productSellerRepo struct {
	db *gorm.DB
}

func NewProductSellerRepo(db *gorm.DB) ProductSellerRepo {
	return &productSellerRepo{db: db}
}

func (p *productSellerRepo) Save(seller *model.ProductSeller) error {
	return p.db.Save(seller).Error
}
