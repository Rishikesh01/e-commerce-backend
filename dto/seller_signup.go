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

func (r *SellerSignup) Error() error {
	if err := r.emptyFieldErr(); err != nil {
		return err
	}

	if err := r.emailFormatErr(); err != nil {
		return err
	}

	if err := r.passwordErr(); err != nil {
		return err
	}

	return nil
}

func (r *SellerSignup) emptyFieldErr() error {
	if r.Name == "" || r.BusinessName == "" || r.Email == "" || r.Password == "" || r.ConfirmPassword == "" {
		return errors.New("field is/are empty")
	}
	return nil
}
func (r *SellerSignup) emailFormatErr() error {
	if _, err := mail.ParseAddress(r.Email); err != nil {
		return err
	}
	return nil
}
func (r *SellerSignup) passwordErr() error {
	if r.Password != r.ConfirmPassword {
		return errors.New("passwords don't match")
	}
	return nil
}
