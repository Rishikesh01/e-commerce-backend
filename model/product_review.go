package model

import (
	"time"
)

type ProductReview struct {
	ID         uint
	ProductID  Product
	UserID     User
	IsVerified bool
	Review     string
	CreatedAT  time.Time
	UpdatedAT  time.Time
}
