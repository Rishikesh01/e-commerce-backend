package dto

type SellerSignup struct {
	Name            string `json:"name"`
	Email           string `json:"email" validate:"required"`
	BusinessName    string `json:"business_name"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password"`
}
