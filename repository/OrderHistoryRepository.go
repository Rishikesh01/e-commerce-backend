package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderHistoryRepository struct {
	Db *gorm.DB
}

func NewOrderHistoryRepository(Db *gorm.DB) *OrderHistoryRepository {
	return &OrderHistoryRepository{Db: Db}
}

func (Order *OrderHistoryRepository) GetUserOrderHistory(UserId uuid.UUID, size uint8) {

}
