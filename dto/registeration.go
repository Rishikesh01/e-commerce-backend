package dto

import (
	"errors"
	"net/mail"
)

type Registration struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (r *Registration) IsAnyFieldEmpty() error {
	if r.Name == "" || r.Email == "" || r.Password == "" || r.ConfirmPassword == "" {
		return errors.New("field is/are empty")
	}
	return nil
}
func (r *Registration) IsValidEmail() error {
	if _, err := mail.ParseAddress(r.Email); err != nil {
		return err
	}
	return nil
}
func (r *Registration) IsPasswordEqual() error {
	if r.Password != r.ConfirmPassword {
		return errors.New("passwords don't match")
	}
	return nil
}
