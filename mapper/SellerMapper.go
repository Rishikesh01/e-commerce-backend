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
