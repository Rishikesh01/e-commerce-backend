package model

import (
	"github.com/google/uuid"
	"time"
)

type ProductReview struct {
	ID         uint
	ProductID  uuid.UUID
	UserID     uuid.UUID
	IsVerified bool
	Review     string
	CreatedAT  time.Time
	UpdatedAT  time.Time
}
