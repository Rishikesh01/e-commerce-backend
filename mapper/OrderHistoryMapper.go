package mapper

import (
	"github.com/Rishikesh01/amazon-clone-backend/dto"
	"github.com/Rishikesh01/amazon-clone-backend/model"
)

func OrderHistoryToOrderHistoryDto(history model.UserOrderHistory) dto.UserOrderHistory {
	var userOrderHistory dto.UserOrderHistory

	return userOrderHistory
}

func OrderHistoryDtoToOrderHistory(history dto.UserOrderHistory) model.UserOrderHistory {
	var userOrderHistory model.UserOrderHistory

	return userOrderHistory
}
