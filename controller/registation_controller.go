package controller

import (
	"github.com/Rishikesh01/amazon-clone-backend/services"
	"net/http"

	"github.com/Rishikesh01/amazon-clone-backend/dto"
	"github.com/gin-gonic/gin"
)

type RegistrationController struct {
	service       services.UserService
	sellerService services.SellerService
}

func NewRegistartionController(service services.UserService, sellerService services.SellerService) *RegistrationController {
	return &RegistrationController{service: service, sellerService: sellerService}
}

func (r *RegistrationController) Signup(ctx *gin.Context) {
	var cred dto.Registration
	if err := ctx.ShouldBindJSON(&cred); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cred.IsAnyFieldEmpty(); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := cred.IsPasswordEqual(); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := cred.IsValidEmail(); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := r.service.Register(cred)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(200)
}

func (r *RegistrationController) SellerSignup(ctx *gin.Context) {
	var cred dto.SellerSignup
	if err := ctx.ShouldBindJSON(&cred); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := r.sellerService.RegisterSeller(cred)
	if err != nil {
		ctx.JSON(400, err)
	}
	ctx.Status(200)
}
