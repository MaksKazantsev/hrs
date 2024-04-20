package models

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	UUID string
	*jwt.RegisteredClaims
}
