package repository

import (
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"github.com/google/uuid"
)

type UserRepo interface {
	Save(user *model.User) error
	FindByID(uuid uuid.UUID) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}
