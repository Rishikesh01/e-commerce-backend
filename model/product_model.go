package model

import "github.com/google/uuid"

type Product struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	PicturePath   string
	Name          string
	Description   string
	HasBasicInfo  bool
	ProductSeller []ProductSeller
	ProductReview []ProductReview
	ProductRating ProductRating
	TrackRating   []TrackRating
}
