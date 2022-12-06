package services

import (
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"github.com/Rishikesh01/amazon-clone-backend/repository"
	"github.com/google/uuid"
)

type BillingService interface {
	CreateBill([]model.Product, uuid.UUID) (*model.Billing, error)
}

func NewBillingService(billingRepo repository.BillingRepo) BillingService {
	return &billingService{billingRepo: billingRepo}
}

type billingService struct {
	billingRepo repository.BillingRepo
}

func (b *billingService) CreateBill(products []model.Product, userID uuid.UUID) (*model.Billing, error) {
	var bill *model.Billing
	for _, val := range products {
		bill.Amount += val.ProductSeller[0].Price
	}
	bill.UserID = userID
	err := b.billingRepo.Save(bill)
	if err != nil {
		return nil, err
	}
	return bill, nil
}
