package mapper

import (
	"github.com/Rishikesh01/amazon-clone-backend/dto"
	"github.com/Rishikesh01/amazon-clone-backend/model"
)

func AddressToAddressDto(Address model.Address) dto.AddressDto {
	var addressDto dto.AddressDto

	return addressDto
}

func AddressDtoToAddress(Address dto.AddressDto) model.Address {
	var address model.Address

	return address
}

func ListOfAddressToListOfAddressDto(Address []model.Address) []dto.AddressDto {
	var addressDto []dto.AddressDto
	for address := range Address {
		addressDto = append(addressDto, AddressToAddressDto(address))
	}
	return addressDto
}
