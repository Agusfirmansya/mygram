package helpers

import "net/mail"

func EmailValidator(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}