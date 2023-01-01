package services

import (
	"errors"
	"github.com/Rishikesh01/amazon-clone-backend/dto"
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"github.com/Rishikesh01/amazon-clone-backend/repository"
	"github.com/Rishikesh01/amazon-clone-backend/util"
	"gorm.io/gorm"
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
	trackRating := &model.TrackRating{ProductID: rating.ID, RatingScore: uint(rating.Rating)}
	trackRating.UserID = rating.UserID
	existingModel, err := u.productRatingRepo.FindByID(rating.ID)
	if err != nil {
		return err
	}
	existingModel.TotalRatingScore = u.getStarRating(
		rating.Rating,
		existingModel.TotalRatingScore,
		int(existingModel.TotalUserRated),
	)
	existingModel.TotalUserRated = existingModel.TotalUserRated + 1
	err = u.productRatingRepo.SaveWithTrackRating(existingModel, trackRating)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) GiveProductReview(review dto.ProductReview) error {
	_, err := u.userRepo.FindByID(review.UserID)
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return err
	}

	err = u.productReviewRepo.Save(&model.ProductReview{
		ProductID: review.ID,
		UserID:    review.UserID,
		Review:    review.Review},
	)
	if err != nil {
		return err
	}

	return nil
}

func (u *userService) getStarRating(newRating int, oldRatingAvg float32, totalRating int) float32 {
	return ((oldRatingAvg * float32(totalRating)) + float32(newRating) + float32(1)) / (float32(totalRating) + float32(1))
}
