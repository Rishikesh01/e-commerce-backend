package controller

import (
	"net/http"

	"github.com/Rishikesh01/amazon-clone-backend/dto"
	"github.com/Rishikesh01/amazon-clone-backend/services"
	"github.com/gin-gonic/gin"
)

type Authication interface {
	Login(ctx *gin.Context)
	Validate(ctx *gin.Context)
}
type JWTAuthController struct {
	LoginService *services.LoginServices
	Auth         services.AuthService
}

func NewJWTAuthController(service *services.LoginServices, auth *services.AuthService) *JWTAuthController {
	return &JWTAuthController{LoginService: service, Auth: *auth}
}

/*
	method used for user Authentication
	endpoint '/login'
*/
func (jwt *JWTAuthController) Login(ctx *gin.Context) {
	var cred dto.Credentials

	if err := ctx.ShouldBindJSON(&cred); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if ok, token := jwt.LoginService.Login(&cred); ok {
		ctx.JSON(http.StatusOK, gin.H{"token:": token})
		return
	}

	ctx.AbortWithStatus(http.StatusUnauthorized)
}

// method used to validate incoming request
func (jwt *JWTAuthController) Validate(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")[7:]

	if _, err := jwt.Auth.ValidateToken(token); err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	ctx.Status(http.StatusOK)
}
