package repository

import (
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"gorm.io/gorm"
)

type TrackProductRatingRepo interface {
	Save(rating *model.TrackRating) error
}

type trackProductRatingRepo struct {
	db *gorm.DB
}

func NewTackProductRatingRepo(db *gorm.DB) TrackProductRatingRepo {
	return &trackProductRatingRepo{db: db}
}

func (t *trackProductRatingRepo) Save(rating *model.TrackRating) error {
	return t.db.Save(rating).Error
}
