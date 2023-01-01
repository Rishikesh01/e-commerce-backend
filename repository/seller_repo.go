package repository

import (
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SellerRepo interface {
	Save(seller *model.Seller) error
	FindByEmail(email string) (*model.Seller, error)
	FindByID(id uuid.UUID) (*model.Seller, error)
}

type sellerRepo struct {
	db *gorm.DB
}

func NewSellerRepo(db *gorm.DB) SellerRepo {
	return &sellerRepo{db: db}
}

func (s *sellerRepo) Save(seller *model.Seller) error {
	return s.db.Save(seller).Error
}

func (s *sellerRepo) FindByEmail(email string) (*model.Seller, error) {
	var seller model.Seller
	if err := s.db.Where("email=?", email).First(&seller).Error; err != nil {
		return nil, err
	}
	return &seller, nil
}

func (s *sellerRepo) FindByID(id uuid.UUID) (*model.Seller, error) {
	var seller model.Seller
	if err := s.db.Where("id=?", id).First(&seller).Error; err != nil {
		return nil, err
	}
	return &seller, nil
}
