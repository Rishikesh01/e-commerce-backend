package repository

import (
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepo interface {
	Save(user *model.User) error
	FindByID(uuid uuid.UUID) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db: db}
}

func (u userRepo) Save(user *model.User) error {
	return u.db.Save(user).Error
}

func (u userRepo) FindByID(uuid uuid.UUID) (*model.User, error) {
	var user model.User
	if err := u.db.Where("id=?", uuid).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u userRepo) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
