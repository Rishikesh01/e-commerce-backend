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

}

func (address *AddressRepository) UpdateUsersAddress(Email string, addr model.Address) {

}
