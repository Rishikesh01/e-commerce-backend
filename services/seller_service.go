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
	AddNewProduct(product dto.AddProduct, sellerID uuid.UUID) error
	AddSellerToExistingProduct(product dto.AddSellerToProduct, sellerID uuid.UUID) error
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

func (p *sellerService) AddNewProduct(product dto.AddProduct, sellerID uuid.UUID) error {

	mProduct := &model.Product{}
	mProduct.Name = product.Name
	mProduct.PicturePath = product.Img
	mProduct.Description = product.Description
	mProduct.ProductSeller = []model.ProductSeller{}
	mProduct.ProductSeller = append(mProduct.ProductSeller, model.ProductSeller{
		SellerID: sellerID,
		Price:    product.Price,
	})

	if err := p.productRepo.Save(mProduct); err != nil {
		return err
	}

	return nil
}

func (p *sellerService) AddSellerToExistingProduct(product dto.AddSellerToProduct, sellerID uuid.UUID) error {
	mProduct, err := p.productRepo.FindByID(product.ID)
	if err != nil {
		return err
	}

	if err := p.productSellerRepo.Save(&model.ProductSeller{
		ProductID: mProduct.ID,
		SellerID:  sellerID,
		Price:     product.Price,
	}); err != nil {
		return err
	}
	return nil
}

func (p *sellerService) UpdateProduct(product model.Product) error {
	return p.productRepo.Save(&product)
}
