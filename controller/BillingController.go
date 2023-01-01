package controller

import (
	"github.com/Rishikesh01/amazon-clone-backend/dto"
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"github.com/Rishikesh01/amazon-clone-backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BillingController struct {
	service services.BillingService
}

func NewBillingController(service services.BillingService) *BillingController {
	return &BillingController{service: service}
}

func (b *BillingController) CreateBill(ctx *gin.Context) {
	var pay dto.BillingDTO
	if err := ctx.ShouldBindJSON(&pay); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	const BEARER_SCHEMA = "Bearer"
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len(BEARER_SCHEMA):]
	if tokenString == "" {
		ctx.Status(403)
		return
	}
	id, _, err := services.GetClaims(tokenString)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	bill, err := b.service.CreateBill(pay.Products, id)

	ctx.JSON(200, model.Billing{ID: bill.ID, UserID: bill.UserID, Amount: bill.Amount})

}
