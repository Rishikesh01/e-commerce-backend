package repository

import "github.com/Rishikesh01/amazon-clone-backend/model"

type SellerRepo interface {
	Save(seller *model.Seller) error
	FindByEmail(email string) (*model.Seller, error)
}
