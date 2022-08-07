package repository

import (
	"log"

	"github.com/Rishikesh01/amazon-clone-backend/dto"
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"github.com/Rishikesh01/amazon-clone-backend/util"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(Db *gorm.DB) *UserRepository {
	return &UserRepository{
		Db: Db,
	}
}

func (u *UserRepository) FindByEmail(email string) *model.User {
	model := &model.User{}
	if err := u.Db.Where("email=?", email).First(&model).Error; err != nil {
		log.Println(err)
		return nil
	}
	return model
}

func (u *UserRepository) AddAddress(email string, address dto.AddressDto) {
	var user model.User
	var deliveryAddress model.Address

	deliveryAddress.AddressLine = address.AddressLine
	deliveryAddress.LandMark = address.LandMark
	deliveryAddress.ZipCode = address.ZipCode
	deliveryAddress.City = address.City
	deliveryAddress.State = address.State
	deliveryAddress.Country = address.Country

	u.Db.Where("email=?", email).First(&user)

	deliveryAddress.UserId = user.Id

	err := u.Db.Model(&user).Association("Addressses").Append([]model.Address{deliveryAddress})

	if err != nil {
		log.Println(err)
	}

}

func (u *UserRepository) GetUserOrderHistory(Email string) {
}

func (u *UserRepository) Save(cred *dto.Credentials) error {
	var user model.User
	user.Name = cred.Email
	user.Email = cred.Email
	pass, err := util.BcryptUtil{}.HashPassword(cred.Password)
	if err != nil {
		return err
	}
	user.Password = pass
	u.Db.Create(&user)
	return nil
}
