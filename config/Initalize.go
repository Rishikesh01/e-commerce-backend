package config

import (
	"github.com/Rishikesh01/amazon-clone-backend/controller"
	"github.com/Rishikesh01/amazon-clone-backend/repository"
	"github.com/Rishikesh01/amazon-clone-backend/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Engine struct{}

func (e *Engine) Run() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))
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
	homeService := services.NewHomePageService(productRepo)

	authController := controller.NewJWTAuthController(authService)
	registrationController := controller.NewRegistrationController(userService, sellerService)
	prodController := controller.NewProductController(productService, sellerService, userService)
	billingController := controller.NewBillingController(billingService)
	homeController := controller.NewHomeController(homeService)

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
	//Image endpoint
	router.Static("/image", "./image")
	//Search endpoint
	router.GET("/search", prodController.SearchForProduct)
	//add New Product Endpoint
	sellerGroup.POST("/seller/product", prodController.AddNewProduct)
	//add Seller to existing product
	sellerGroup.POST("/seller/existing/product", prodController.AddToExistingProduct)
	//billing endpoint
	sGroup.POST("/user/bill", billingController.CreateBill)

	router.GET("/home", homeController.Home)
	sGroup.POST("/give/rating", prodController.RateProduct)
	sGroup.POST("/give/comment", prodController.LeaveComment)

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
