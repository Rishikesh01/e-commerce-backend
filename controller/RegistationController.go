package controller

import (
	"net/http"

	"github.com/Rishikesh01/amazon-clone-backend/dto"
	"github.com/Rishikesh01/amazon-clone-backend/repository"
	"github.com/gin-gonic/gin"
)

type RegistartionController struct{
 UserRepo  *repository.UserRepository
}

func NewRegistartionController(repo *repository.UserRepository) *RegistartionController{
  return &RegistartionController{UserRepo:repo}
}

func (r *RegistartionController) Signup(ctx *gin.Context){
	var cred dto.Credentials
	if err:=ctx.ShouldBindJSON(&cred);err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"error":err})
	return
	}
	r.UserRepo.Save(&cred)
	ctx.JSON(200,"test")
}
