package dto

type Registration struct{
	Email string	`json:"email"`
	Password string	`json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}