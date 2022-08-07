package mapper

import (
	"github.com/Rishikesh01/amazon-clone-backend/dto"
	"github.com/Rishikesh01/amazon-clone-backend/model"
)

func ProductToProductDto(product model.Product) dto.Product {
	var productDto dto.Product

	return productDto
}

func ProductDtoToProduct(product dto.Product) model.Product {
	var productModel model.Product

	return productModel
}
