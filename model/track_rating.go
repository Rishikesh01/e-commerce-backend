package model

import "github.com/google/uuid"

type TrackRating struct {
	ID          uint
	ProductID   uuid.UUID
	UserID      []User
	RatingScore uint
}
