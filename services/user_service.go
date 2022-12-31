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
	RateProduct(rating dto.ProductRatingByUser) error
	GiveProductReview(review dto.ProductReview) error
}

type userService struct {
	productRatingRepo repository.ProductRatingRepo
	productReviewRepo repository.ProductReviewRepo
	trackRatingRepo   repository.TrackProductRatingRepo
	userRepo          repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo, productRatingRepo repository.ProductRatingRepo, productReviewRepo repository.ProductReviewRepo, trackRatingRepo repository.TrackProductRatingRepo) UserService {
	return &userService{userRepo: userRepo, productReviewRepo: productReviewRepo, productRatingRepo: productRatingRepo, trackRatingRepo: trackRatingRepo}
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

func (u *userService) RateProduct(rating dto.ProductRatingByUser) error {
	panic("implement")
	return nil
}

func (u *userService) GiveProductReview(review dto.ProductReview) error {
	panic("implement")
	return nil
}
