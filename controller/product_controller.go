package controller

import (
	"fmt"
	"github.com/Rishikesh01/amazon-clone-backend/dto"
	"github.com/Rishikesh01/amazon-clone-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"os"
)

type ProductController struct {
	productService services.ProductService
}

func NewProductController(service services.ProductService) ProductController {
	return ProductController{
		productService: service,
	}
}

func (p *ProductController) SearchForProduct(ctx *gin.Context) {
	search := ctx.Param("search")
	products, err := p.productService.Search(search)
	var productDtos []dto.ProductSearch
	for i := 0; i < len(products); i++ {
		val := dto.ProductSearch{
			ID:          products[i].ID,
			Name:        products[i].Name,
			Description: products[i].Description,
			Sellers:     products[i].ProductSeller,
		}
		productDtos = append(productDtos, val)
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			ctx.JSON(404, err)
			return
		}
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(200, products)
}

func (p *ProductController) AddNewProductPicture(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	dir, err := os.UserHomeDir()
	if err != nil {
		return
	}
	fileName := dir + "images" + file.Filename + uuid.New().String() + ".jpg"
	err = ctx.SaveUploadedFile(file, fileName)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	const BEARER_SCHEMA = "Bearer"
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len(BEARER_SCHEMA)+1:]
	id, _, _, err := services.GetSellerClaims(tokenString)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ID, err := p.productService.AddNewProductImage(fileName, id)

	ctx.JSON(http.StatusOK, fmt.Sprintf("ID:%s", ID.String()))
}

// adds product
func (p *ProductController) AddNewProduct(ctx *gin.Context) {
	var prod dto.Product
	if err := ctx.ShouldBindJSON(&prod); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	err := p.productService.AddNewProduct(prod)
	if err != nil {
		ctx.JSON(400, err)
		return
	}
}
