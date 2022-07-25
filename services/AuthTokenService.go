package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	GenerateToken(email string, isUser bool) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
type authCustomClaims struct {
	Email  string `json:"email"`
	IsUser bool   `json:"isUser"`
	jwt.RegisteredClaims
}

type AuthService struct {
	issuer    string
	secretKey string
}

func NewAuthService() *AuthService {
	return &AuthService{
		secretKey: "test",
		issuer:    "amazon-backend-clone",
	}
}

func (service *AuthService) GenerateToken(email string, isUser bool) (string, error) {
	claims := &authCustomClaims{
		email,
		isUser,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15).UTC()),
			Issuer:    service.issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return signedToken, nil
}

func (service *AuthService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})
}
