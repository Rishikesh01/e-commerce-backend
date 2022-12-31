package model

type ProductRating struct {
	ID               uint
	TrackRatingID    TrackRating
	TotalRatingScore uint
	TotalUserRated   uint
}
