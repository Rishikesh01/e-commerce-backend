package repository

import (
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AddressRepository struct {
	Db *gorm.DB
}

func NewAddressRepository(Db *gorm.DB) *AddressRepository {
	return &AddressRepository{Db: Db}
}

func (address *AddressRepository) UpdateSellersAddress(id uuid.UUID, addr model.Address) {
	address.Db.Model(&model.Address{}).Where("seller_id=?", id.String()).Save(&addr)
}

func (address *AddressRepository) UpdateUsersAddress(id uuid.UUID, userId uuid.UUID, addr model.Address) {
	address.Db.Model(&model.Address{}).Where("id=? AND user_id=?", id.String(), userId.String()).Save(addr)
}
