package repository

import (
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductSellerRepo interface {
	FindByCompositeID(ProductId, SellerID uuid.UUID) (*model.ProductSeller, error)
	Save(*model.ProductSeller) error
}

type productSellerRepo struct {
	db *gorm.DB
}

func NewProductSellerRepo(db *gorm.DB) ProductSellerRepo {
	return &productSellerRepo{db: db}
}

func (p *productSellerRepo) FindByCompositeID(ProductId, SellerID uuid.UUID) (*model.ProductSeller, error) {
	var ps model.ProductSeller

	if err := p.db.Where("product_id=? and seller_id=?", ProductId, SellerID).First(&ps).Error; err != nil {
		return nil, err
	}

	return &ps, nil
}

func (p *productSellerRepo) Save(seller *model.ProductSeller) error {
	return p.db.Save(seller).Error
}
