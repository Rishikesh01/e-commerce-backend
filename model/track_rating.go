package model

type TrackRating struct {
	ID          uint
	ProductID   Product
	UserID      []User
	RatingScore uint
}
