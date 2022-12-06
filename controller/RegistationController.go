package controller

import (
	"github.com/Rishikesh01/amazon-clone-backend/services"
	"net/http"

	"github.com/Rishikesh01/amazon-clone-backend/dto"
	"github.com/gin-gonic/gin"
)

type RegistrationController struct {
	service services.UserService
}

func NewRegistartionController(service services.UserService) *RegistrationController {
	return &RegistrationController{service: service}
}

func (r *RegistrationController) Signup(ctx *gin.Context) {
	var cred dto.Registration
	if err := ctx.ShouldBindJSON(&cred); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := r.service.Register(cred)
	if err != nil {
		ctx.JSON(400, err)
	}
	ctx.JSON(200, "success")
}
