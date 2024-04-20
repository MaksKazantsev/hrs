package db

import (
	"context"
	"github.com/alserov/hrs/auth/internal/models"
)

type Repository interface {
	SignUp(ctx context.Context, req models.RegReq) error
	SignIn(ctx context.Context, email string) (LoginInfo, error)
}

type LoginInfo struct {
	UUID     string `json:"uuid"`
	Email    string
	Password string
	Username string
}
