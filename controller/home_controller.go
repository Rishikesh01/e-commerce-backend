package controller

import (
	"github.com/Rishikesh01/amazon-clone-backend/services"
	"github.com/gin-gonic/gin"
)

type HomeController struct {
	homePageService services.HomePageService
}

func NewHomeController(service services.HomePageService) *HomeController {
	return &HomeController{homePageService: service}
}

func (h *HomeController) Home(ctx *gin.Context) {
	data, err := h.homePageService.ShowItems()
	if err != nil {
		ctx.Status(500)
		return
	}
	ctx.JSON(200, data)
}
