package controller

import (
	"net/http"

	"github.com/Rishikesh01/amazon-clone-backend/dto"
	"github.com/Rishikesh01/amazon-clone-backend/services"
	"github.com/gin-gonic/gin"
)

type JWTAuthController struct {
	userAuth services.AuthService
}

func NewJWTAuthController(auth services.AuthService) *JWTAuthController {
	return &JWTAuthController{userAuth: auth}
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

	if token, err := jwt.userAuth.AuthUser(cred, "user"); err != nil {
		ctx.Status(http.StatusUnauthorized)

	} else {
		ctx.JSON(http.StatusOK, gin.H{"token:": token})
		return
	}

}

func (jwt *JWTAuthController) SellerLogin(ctx *gin.Context) {
	var cred dto.Credentials

	if err := ctx.ShouldBindJSON(&cred); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if token, err := jwt.userAuth.AuthUser(cred, "seller"); err != nil {
		ctx.Status(http.StatusUnauthorized)

	} else {
		ctx.JSON(http.StatusOK, gin.H{"token:": token})
		return
	}
}
