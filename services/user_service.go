package services

import (
	"errors"
	"github.com/Rishikesh01/amazon-clone-backend/dto"
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"github.com/Rishikesh01/amazon-clone-backend/repository"
	"github.com/Rishikesh01/amazon-clone-backend/util"
)

type UserService interface {
	Register(registration dto.Registration) error
	RegisterSeller(registration dto.SellerSignup) error
}

type userService struct {
	userRepo   repository.UserRepo
	sellerRepo repository.SellerRepo
}

func NewUserService(userRepo repository.UserRepo, sellerRepo repository.SellerRepo) UserService {
	return &userService{userRepo: userRepo, sellerRepo: sellerRepo}
}

func (u *userService) Register(registration dto.Registration) error {
	if registration.Password != registration.ConfirmPassword {
		return errors.New("passwords don't match")
	}
	utility := util.BcryptUtil{}
	password, err := utility.HashPassword(registration.ConfirmPassword)
	if err != nil {
		return err
	}
	user := &model.User{Name: registration.Name, Email: registration.Email, Password: password}
	if err = u.userRepo.Save(user); err != nil {
		return err
	}
	return nil
}

func (u *userService) RegisterSeller(registration dto.SellerSignup) error {
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
