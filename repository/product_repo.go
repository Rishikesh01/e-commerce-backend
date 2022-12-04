package repository

import "github.com/Rishikesh01/amazon-clone-backend/model"

type ProductRepo interface {
	Save(product *model.Product) error
	FindByLikeName(name string) ([]model.Product, error)
}
