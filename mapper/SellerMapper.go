package mapper

import (
	"github.com/Rishikesh01/amazon-clone-backend/dto"
	"github.com/Rishikesh01/amazon-clone-backend/model"
)

func SellerToSellerDto(seller model.Seller) dto.SellerDto {
	var SellerDto dto.SellerDto

	return SellerDto
}

func SellerDtoToSeller(seller dto.SellerDto) model.Seller {
	var sellerModel model.Seller

	return sellerModel
}

func ListOfSellerDtoToSeller(seller []dto.SellerDto) []model.Seller {
	var sellerModels []model.Seller

	for _, sellerDto := range seller {
		sellerModels = append(sellerModels, SellerDtoToSeller(sellerDto))
	}

	return sellerModels
}
