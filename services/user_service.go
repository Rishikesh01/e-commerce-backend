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
}

type userService struct {
	userRepo repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo) UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) Register(registration dto.Registration) error {
	if registration.Password != registration.ConfirmPassword {
		return errors.New("wrong Pass")
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
