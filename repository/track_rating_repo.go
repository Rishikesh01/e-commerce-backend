package repository

import "gorm.io/gorm"

type TrackProductRatingRepo interface {
}

type trackProductRatingRepo struct {
	db *gorm.DB
}

func NewTackProductRatingRepo(db *gorm.DB) TrackProductRatingRepo {
	return &trackProductRatingRepo{db: db}
}
