package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(pass string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", NewError(ErrInternal, err.Error())
	}
	return string(b), nil
}

func CompareHash(hash, pass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass)); err != nil {
		return NewError(ErrBadRequest, "invalid password")
	}
	return nil
}
