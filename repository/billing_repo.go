package repository

import (
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"github.com/google/uuid"
)

type BillingRepo interface {
	Save(billing *model.Billing) error
	FindByUserID(uuid uuid.UUID) (*model.Billing, error)
}
