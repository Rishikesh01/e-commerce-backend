package config

import (
	"github.com/Rishikesh01/amazon-clone-backend/controller"
	"github.com/Rishikesh01/amazon-clone-backend/repository"
	"github.com/Rishikesh01/amazon-clone-backend/services"
	"github.com/gin-gonic/gin"
	"log"
)

type Engine struct{}

func (e *Engine) Run() {
	router := gin.Default()
	start(router)
}

func start(router *gin.Engine) {
	db := repository.Init()

	userRepo := repository.NewUserRepo(db)
	//productSellerRepo := repository.NewProductSellerRepo(db)
	//sellerRepo := repository.NewSellerRepo(db)
	//billingRepo := repository.NewBillingRepo(db)
	//productRepo := repository.NewProductRepo(db)

	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userRepo)
	//productService := services.NewProductService(productRepo,sellerRepo,productSellerRepo)
	//billingService := services.NewBillingService(billingRepo)

	authController := controller.NewJWTAuthController(authService)
	registrationController := controller.NewRegistartionController(userService)

	router.POST("/register", registrationController.Signup)
	router.POST("/login", authController.Login)

	err := router.Run()
	if err != nil {
		log.Println(err)
	}
}
