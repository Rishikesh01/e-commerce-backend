package repository

import (
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BillingRepo interface {
	Save(billing *model.Billing) error
	FindByID(uuid uuid.UUID) (*model.Billing, error)
}

type billingRepo struct {
	db *gorm.DB
}

func NewBillingRepo(db *gorm.DB) BillingRepo {
	return &billingRepo{db: db}
}

func (b *billingRepo) Save(billing *model.Billing) error {
	return b.db.Save(billing).Error
}

func (b *billingRepo) FindByID(userid uuid.UUID) (*model.Billing, error) {
	var bill model.Billing
	if err := b.db.Where("user_id = ?", userid).First(&bill).Error; err != nil {
		return nil, err
	}

	return &bill, nil
}
