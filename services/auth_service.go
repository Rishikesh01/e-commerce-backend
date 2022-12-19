package services

import (
	"errors"
	"github.com/Rishikesh01/amazon-clone-backend/dto"
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"github.com/Rishikesh01/amazon-clone-backend/repository"
	"github.com/Rishikesh01/amazon-clone-backend/util"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"log"
	"time"
)

type AuthService interface {
	AuthUser(credentials dto.Credentials, userType string) (string, error)
	ValidateToken(token string, tokenType string) error
}

func NewAuthService(userRepo repository.UserRepo, sellerRepo repository.SellerRepo) AuthService {
	return &jwtAuthService{userRepo: userRepo, sellerRepo: sellerRepo}
}

type customSellerJWT struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	BusinessName string    `json:"business_name"`
	IsPrime      bool      `json:"is_prime"`
	jwt.RegisteredClaims
}

type customUserJWT struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	jwt.RegisteredClaims
}

type jwtAuthService struct {
	userRepo   repository.UserRepo
	sellerRepo repository.SellerRepo
}

func (j *jwtAuthService) AuthUser(credentials dto.Credentials, userType string) (string, error) {
	if userType == "user" {
		user, err := j.loadUser(credentials.Email)
		if err != nil {
			return "", nil
		}
		utility := util.BcryptUtil{}
		if utility.CheckPasswordHash(credentials.Password, user.Password) {
			token, err := j.createUserToken(credentials.Email, user.Id)
			if err != nil {
				return "", err
			}
			return token, err
		}

		return "", errors.New("wrong username/password")
	}

	user, err := j.loadSeller(credentials.Email)
	if err != nil {
		return "", nil
	}
	utility := util.BcryptUtil{}
	if utility.CheckPasswordHash(credentials.Password, user.Password) {
		token, err := j.createSellerToken(credentials.Email, user.BusinessName, user.ID)
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

func (j *jwtAuthService) loadSeller(email string) (*model.Seller, error) {
	user, err := j.sellerRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, err
}
func (j *jwtAuthService) createUserToken(email string, id uuid.UUID) (string, error) {
	claim := customUserJWT{
		id,
		email,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "e-commerce",
			Subject:   "user:" + email,
			ID:        id.String(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claim)

	signedToken, err := token.SignedString([]byte("demo"))
	if err != nil {
		log.Println(err)
		return "", err
	}
	return signedToken, nil
}

func (j *jwtAuthService) createSellerToken(email string, businessName string, id uuid.UUID) (string, error) {
	claim := customSellerJWT{
		id,
		email,
		businessName,
		true,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "e-commerce",
			Subject:   "seller:" + email,
			ID:        id.String(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claim)

	signedToken, err := token.SignedString([]byte("demo"))
	if err != nil {
		log.Println(err)
		return "", err
	}
	return signedToken, nil
}

func (j *jwtAuthService) ValidateToken(token string, tokenType string) error {
	if tokenType == "user" {
		t, err := jwt.ParseWithClaims(
			token,
			&customUserJWT{},
			func(token *jwt.Token) (interface{}, error) {
				return "demo", nil
			},
		)
		if err != nil {
			return nil
		}
		claims, ok := t.Claims.(*customUserJWT)
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

	t, err := jwt.ParseWithClaims(
		token,
		&customSellerJWT{},
		func(token *jwt.Token) (interface{}, error) {
			return "demo", nil
		},
	)
	if err != nil {
		return nil
	}
	claims, ok := t.Claims.(*customSellerJWT)
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

func GetClaims(token string) (uuid.UUID, string, error) {
	t, err := jwt.ParseWithClaims(
		token,
		&customUserJWT{},
		func(token *jwt.Token) (interface{}, error) {
			return "demo", nil
		},
	)
	if err != nil {
		return uuid.UUID{}, "", nil
	}
	claims, _ := t.Claims.(*customUserJWT)
	return claims.ID, claims.Email, nil
}

func GetSellerClaims(token string) (uuid.UUID, string, string, error) {
	t, err := jwt.ParseWithClaims(
		token,
		&customSellerJWT{},
		func(token *jwt.Token) (interface{}, error) {
			return "demo", nil
		},
	)
	if err != nil {
		return uuid.UUID{}, "", "", nil
	}
	claims, _ := t.Claims.(*customSellerJWT)
	return claims.ID, claims.Email, claims.BusinessName, nil

}
