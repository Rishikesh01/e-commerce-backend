package repository

import (
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SellerRepository struct {
	Db *gorm.DB
}

func NewSeller(Db *gorm.DB) *SellerRepository {
	return &SellerRepository{Db: Db}
}

func (s *SellerRepository) Save(seller model.Seller) {
	s.Db.Create(&seller)
}

func (s *SellerRepository) AddSellersProduct(SellerId uuid.UUID) {}
