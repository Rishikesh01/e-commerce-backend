package repository

import (
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"github.com/google/uuid"
)

type UserOrderHistoryRepo interface {
	Save(order *model.UserOrderHistory) error
	FindByBillingID(uuid uuid.UUID) ([]model.UserOrderHistory, error)
}
