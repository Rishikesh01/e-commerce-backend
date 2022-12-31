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
	trackRatingRepo := repository.NewTackProductRatingRepo(db)
	productRatingRepo := repository.NewProductRatingRepo(db)
	productReviewRepo := repository.NewProductReviewRepo(db)

	userService := services.NewUserService(userRepo, productRatingRepo, productReviewRepo, trackRatingRepo)
	sellerService := services.NewSellerService(productRepo, productSellerRepo, sellerRepo)
	authService := services.NewAuthService(userRepo, sellerRepo)
	productService := services.NewProductService(productRepo, sellerRepo, productSellerRepo)
	billingService := services.NewBillingService(billingRepo)

	authController := controller.NewJWTAuthController(authService)
	registrationController := controller.NewRegistartionController(userService, sellerService)
	prodController := controller.NewProductController(productService, sellerService)
	billingController := controller.NewBillingController(billingService)

	//Groups
	sGroup := router.Group("/s/user").Use(authMiddle(authService))
	sellerGroup := router.Group("/s/seller").Use(sellerAuthMiddleWare(authService))

	//seller login
	router.POST("/seller/login", authController.SellerLogin)
	//user signup endpoint
	router.POST("/register", registrationController.Signup)
	//seller signup endpoint
	router.POST("/seller/register", registrationController.SellerSignup)
	//user login endpoint
	router.POST("/login", authController.Login)
	//search endpoint
	router.GET("/search", prodController.SearchForProduct)
	//add New Product Endpoint
	sellerGroup.POST("/seller/product/image", prodController.AddNewProductPicture)
	sellerGroup.POST("/seller/product", prodController.AddNewProduct)
	//billing endpoint
	sGroup.POST("/user/bill", billingController.CreateBill)

	//WIP
	router.GET("/home")
	router.GET("/seller/product/rating")
	sGroup.POST("/give/rating")
	sGroup.POST("/give/comment")

	//TODO
	// Add endpoints to show random products in home screen
	// ADD rating support
	// Add comment support
	// Add Image support

	err := router.Run()
	if err != nil {
		log.Println(err)
	}
}

func sellerAuthMiddleWare(service services.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := ctx.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA)+1:]
		err := service.ValidateToken(tokenString, "seller")
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Next()
	}
}
func authMiddle(service services.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := ctx.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA)+1:]
		err := service.ValidateToken(tokenString, "user")
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Next()
	}
}
