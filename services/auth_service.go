package services

import (
	"errors"
	"github.com/Rishikesh01/amazon-clone-backend/dto"
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"github.com/Rishikesh01/amazon-clone-backend/repository"
	"github.com/Rishikesh01/amazon-clone-backend/util"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
)

type AuthService interface {
	AuthUser(credentials dto.Credentials) (string, error)
	ValidateToken(token string) error
}

func NewAuthService(userRepo repository.UserRepo) AuthService {
	return &jwtAuthService{userRepo: userRepo}
}

type customJWTToken struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	jwt.RegisteredClaims
}

type jwtAuthService struct {
	userRepo repository.UserRepo
}

func (j *jwtAuthService) AuthUser(credentials dto.Credentials) (string, error) {
	user, err := j.loadUser(credentials.Email)
	if err != nil {
		return "", nil
	}
	utility := util.BcryptUtil{}
	if utility.CheckPasswordHash(credentials.Password, user.Password) {
		token, err := j.createToken(credentials.Email, user.Id)
		if err != nil {
			return "", err
		}
		return token, err
	}

	return "", errors.New("wrong username/password")

}

func (j *jwtAuthService) loadUser(email string) (*model.User, error) {
	user, err := j.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (j *jwtAuthService) createToken(email string, id uuid.UUID) (string, error) {
	claim := customJWTToken{
		ID:    id,
		Email: email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	signedToken, err := token.SignedString("demo")
	return signedToken, err
}

func (j *jwtAuthService) ValidateToken(token string) error {
	t, err := jwt.ParseWithClaims(
		token,
		&customJWTToken{},
		func(token *jwt.Token) (interface{}, error) {
			return "demo", nil
		},
	)
	if err != nil {
		return nil
	}
	claims, ok := t.Claims.(*customJWTToken)
	if !ok {
		err = errors.New("couldn't parse claims")
		return err
	}
	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return err
	}
	return nil
}
