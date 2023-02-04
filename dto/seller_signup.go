package dto

import (
	"errors"
	"net/mail"
)

type SellerSignup struct {
	Name            string `json:"name"`
	Email           string `json:"email" validate:"required"`
	BusinessName    string `json:"business_name"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password"`
}

func (r *SellerSignup) IsAnyFieldEmpty() error {
	if r.Name == "" || r.BusinessName == "" || r.Email == "" || r.Password == "" || r.ConfirmPassword == "" {
		return errors.New("field is/are empty")
	}
	return nil
}
func (r *SellerSignup) IsValidEmail() error {
	if _, err := mail.ParseAddress(r.Email); err != nil {
		return err
	}
	return nil
}
func (r *SellerSignup) IsPasswordEqual() error {
	if r.Password != r.ConfirmPassword {
		return errors.New("passwords don't match")
	}
	return nil
}
