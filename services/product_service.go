package services

import (
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"github.com/Rishikesh01/amazon-clone-backend/repository"
)

type ProductService interface {
	Search(name string) ([]model.Product, error)
}

type productService struct {
	productSellerRepo repository.ProductSellerRepo
	sellerRepo        repository.SellerRepo
	productRepo       repository.ProductRepo
}

func NewProductService(productRepo repository.ProductRepo, sellerRepo repository.SellerRepo, productSellerRepo repository.ProductSellerRepo) ProductService {
	return &productService{productRepo: productRepo, sellerRepo: sellerRepo, productSellerRepo: productSellerRepo}
}

func (p *productService) Search(name string) ([]model.Product, error) {
	result, err := p.productRepo.FindByLikeName(name)
	if err != nil {
		return nil, err
	}

	return result, nil
}
