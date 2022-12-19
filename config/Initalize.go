package config

import (
	"github.com/Rishikesh01/amazon-clone-backend/controller"
	"github.com/Rishikesh01/amazon-clone-backend/repository"
	"github.com/Rishikesh01/amazon-clone-backend/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Engine struct{}

func (e *Engine) Run() {
	router := gin.Default()
	start(router)
}

func start(router *gin.Engine) {
	db := repository.Init()

	userRepo := repository.NewUserRepo(db)
	productSellerRepo := repository.NewProductSellerRepo(db)
	sellerRepo := repository.NewSellerRepo(db)
	billingRepo := repository.NewBillingRepo(db)
	productRepo := repository.NewProductRepo(db)

	userService := services.NewUserService(userRepo, sellerRepo)
	authService := services.NewAuthService(userRepo)
	productService := services.NewProductService(productRepo, sellerRepo, productSellerRepo)
	billingService := services.NewBillingService(billingRepo)

	authController := controller.NewJWTAuthController(authService)
	registrationController := controller.NewRegistartionController(userService)
	prodController := controller.NewProductController(productService)
	billingController := controller.NewBillingController(billingService)

	router.POST("/register", registrationController.Signup)
	router.POST("/login", authController.Login)
	router.GET("/search", prodController.SearchForProduct)

	sGroup := router.Group("/s/user")
	sGroup.Use(authMiddle(authService))
	sGroup.POST("/product", prodController.AddNewProduct)
	sGroup.POST("/user/bill", billingController.CreateBill)

	//TODO
	// ADD user rating
	// Add comment support
	// Add Image support

	err := router.Run()
	if err != nil {
		log.Println(err)
	}
}

func authMiddle(service services.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := ctx.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		err := service.ValidateToken(tokenString)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Next()
	}
}
