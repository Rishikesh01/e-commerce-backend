package services

import (
	"log"

	"github.com/Rishikesh01/amazon-clone-backend/dto"
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"github.com/Rishikesh01/amazon-clone-backend/repository"
	"github.com/Rishikesh01/amazon-clone-backend/util"
)

type LoginServices struct {
	repo    *repository.UserRepository
	utility util.BcryptUtil
	auth    *AuthService
}

func NewLoginService(repository *repository.UserRepository, util *util.BcryptUtil,
	jwt *AuthService) *LoginServices {
	return &LoginServices{repo: repository, utility: *util, auth: jwt}
}

func (l *LoginServices) Login(creds *dto.Credentials) (bool, string) {
	email := creds.Email
	password := creds.Password

	var userModel *model.User
	if userModel = l.repo.FindByEmail(email); userModel == nil {
		return false, ""
	}

	if l.utility.CheckPasswordHash(password, userModel.Password) {
		token, err := l.auth.GenerateToken(email, true)
		if err != nil {
			log.Println(err)
			return false, ""
		}
		return true, token
	}

	return false, ""
}
