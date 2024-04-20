package utils

import (
	"github.com/alserov/hrs/auth/internal/models"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

const (
	ENV_SECRET_KEY = "SECRET_KEY"
)

func NewToken(uuid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, models.Claims{
		UUID: uuid,
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 21)),
		},
	})

	tokenString, err := token.SignedString([]byte(os.Getenv(ENV_SECRET_KEY)))
	if err != nil {
		return "", NewError(ErrInternal, err.Error())
	}

	return tokenString, err
}

func ParseToken(token string) (string, error) {
	claims := models.Claims{}
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv(ENV_SECRET_KEY)), nil
	})
	if err != nil {
		return "", NewError(ErrBadRequest, "wrong token provided")
	}
	return claims.UUID, nil
}
