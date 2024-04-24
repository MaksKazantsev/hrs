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
	Verificate(ctx context.Context, code, email string) error
	ResetPass(ctx context.Context, uuid string, password string) error
	RecoverPass(ctx context.Context, req models.RecoverReq) error
	GetVerif(ctx context.Context, email string) (models.VerInfo, error)
}

type Actions interface {
	GetUserInfoByID(ctx context.Context, uuid string) (models.UserInfo, error)
	GetUserInfoByEmail(ctx context.Context, email string) (models.UserInfo, error)
	GetUserPassword(ctx context.Context, uuid string) (string, error)
	SaveVerif(ctx context.Context, info models.VerInfo) error
}

type LoginInfo struct {
	UUID     string `json:"uuid"`
	Email    string
	Password string
	Username string
}
