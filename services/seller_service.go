package services

import (
	"errors"
	"github.com/Rishikesh01/amazon-clone-backend/dto"
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"github.com/Rishikesh01/amazon-clone-backend/repository"
	"github.com/Rishikesh01/amazon-clone-backend/util"
	"github.com/google/uuid"
)

type SellerService interface {
	RegisterSeller(registration dto.SellerSignup) error
	AddNewProduct(product dto.Product) error
	AddNewProductImage(ImagePath string, sellerID uuid.UUID) (uuid.UUID, error)
	UpdateProduct(product model.Product) error
}

type sellerService struct {
	productSellerRepo repository.ProductSellerRepo
	sellerRepo        repository.SellerRepo
	productRepo       repository.ProductRepo
}

func NewSellerService(productRepo repository.ProductRepo, productSellerRepo repository.ProductSellerRepo, sellerRepo repository.SellerRepo) SellerService {
	return &sellerService{productSellerRepo: productSellerRepo, productRepo: productRepo, sellerRepo: sellerRepo}
}

func (u *sellerService) RegisterSeller(registration dto.SellerSignup) error {
	if registration.Password != registration.ConfirmPassword {
		return errors.New("passwords don't match")
	}
	utility := util.BcryptUtil{}
	password, err := utility.HashPassword(registration.ConfirmPassword)
	if err != nil {
		return err
	}
	if err := u.sellerRepo.Save(&model.Seller{Name: registration.Name, Email: registration.Email, BusinessName: registration.BusinessName, Password: password}); err != nil {
		return err
	}
	return nil
}

func (p *sellerService) AddNewProductImage(ImagePath string, sellerID uuid.UUID) (uuid.UUID, error) {
	mProduct := &model.Product{PicturePath: ImagePath}
	mProduct.ProductSeller = append(mProduct.ProductSeller, model.ProductSeller{SellerID: sellerID})

	if err := p.productRepo.Save(mProduct); err != nil {
		return uuid.UUID{}, err
	}

	return mProduct.ID, nil
}

func (p *sellerService) AddNewProduct(product dto.Product) error {
	mProduct := &model.Product{
		ID:            product.ID,
		Name:          product.Name,
		Description:   product.Description,
		ProductSeller: []model.ProductSeller{}}
	mProduct.HasBasicInfo = true
	err := p.productRepo.Save(mProduct)
	if err != nil {
		return err
	}

	return nil
}

func (p *sellerService) UpdateProduct(product model.Product) error {
	return p.productRepo.Save(&product)
}
