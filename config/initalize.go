package config

import (
	"log"

	"github.com/Rishikesh01/amazon-clone-backend/controller"
	database "github.com/Rishikesh01/amazon-clone-backend/model"
	"github.com/Rishikesh01/amazon-clone-backend/repository"
	"github.com/Rishikesh01/amazon-clone-backend/services"
	"github.com/Rishikesh01/amazon-clone-backend/util"
	"github.com/gin-gonic/gin"
)

type Engine struct{}

func (e *Engine) Run() {
	router := gin.Default()
	start(router)
}

func start(router *gin.Engine) {
	util := &util.BcryptUtil{}
	DB := database.Init()
	userRepository := repository.NewUserRepository(DB)
	jwtService := services.NewAuthService()
	loginService := services.NewLoginService(userRepository, util, jwtService)
	authController := controller.NewJWTAuthController(loginService, jwtService)
	reg := controller.NewRegistartionController(userRepository)

	router.POST("/login", authController.Login)
	router.POST("/register", reg.Signup)
	groups := router.Group("/").Use(authController.Validate)
	groups.GET("/order/order-history")
	groups.GET("/product-data")
	groups.POST("/order/order-placed")
	err := router.Run()
	if err != nil {
		log.Println(err)
	}
}
