package services

import (
	"github.com/Rishikesh01/amazon-clone-backend/dto"
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"github.com/Rishikesh01/amazon-clone-backend/repository"
	"github.com/google/uuid"
	"log"
)

type ProductService interface {
	AddNewProduct(product dto.Product, sellerID uuid.UUID) error
	Search(name string) ([]model.Product, error)
	Update(product model.Product) error
}

type productService struct {
	productSellerRepo repository.ProductSellerRepo
	sellerRepo        repository.SellerRepo
	productRepo       repository.ProductRepo
}

func NewProductService(productRepo repository.ProductRepo, sellerRepo repository.SellerRepo, productSellerRepo repository.ProductSellerRepo) ProductService {
	return &productService{productRepo: productRepo, sellerRepo: sellerRepo, productSellerRepo: productSellerRepo}
}

func (p *productService) AddNewProduct(product dto.Product, sellerID uuid.UUID) error {
	_, err := p.sellerRepo.FindByID(sellerID)
	if err != nil {
		log.Println(err)
		return err
	}
	mProduct := &model.Product{
		Name:          product.Name,
		Description:   product.Description,
		ProductSeller: []model.ProductSeller{}}

	mProduct.ProductSeller = append(mProduct.ProductSeller, model.ProductSeller{SellerID: sellerID, Price: product.Price})

	err = p.productRepo.Save(mProduct)
	if err != nil {
		return err
	}

	return nil
}

func (p *productService) Search(name string) ([]model.Product, error) {
	result, err := p.productRepo.FindByLikeName(name)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *productService) Update(product model.Product) error {
	return p.productRepo.Save(&product)
}
