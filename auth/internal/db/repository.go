package db

import (
	"context"
	"github.com/alserov/hrs/auth/internal/models"
)

type Repository interface {
	Auth
	Actions
}

type Auth interface {
	SignUp(ctx context.Context, req models.RegReq) error
	SignIn(ctx context.Context, email string) (LoginInfo, error)
	ResetPass(ctx context.Context, uuid string, password string) error
}

type Actions interface {
	GetUserInfo(ctx context.Context, uuid string) (models.User, error)
	GetUserPassword(ctx context.Context, uuid string) (string, error)
}

type LoginInfo struct {
	UUID     string `json:"uuid"`
	Email    string
	Password string
	Username string
}
