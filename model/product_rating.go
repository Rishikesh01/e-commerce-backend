package model

import "github.com/google/uuid"

type ProductRating struct {
	ID               uint
	ProductID        uuid.UUID
	TotalRatingScore float32
	TotalUserRated   uint
}
