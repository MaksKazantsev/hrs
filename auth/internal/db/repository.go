package db

import (
	"context"
	"github.com/alserov/hrs/auth/internal/models"
)

type Repository interface {
	Auth
	Actions
	Verification
}

type Auth interface {
	SignUp(ctx context.Context, req models.RegReq) error
	SignIn(ctx context.Context, email string) (LoginInfo, error)
	ResetPass(ctx context.Context, uuid string, password string) error
	RecoverPass(ctx context.Context, req models.RecoverReq) error
}

type Verification interface {
	VerificateRecover(ctx context.Context, code, email, password string) error
	Verificate(ctx context.Context, code, email string) error
	GetVerification(ctx context.Context, email string) (models.VerInfo, error)
	GetRecover(ctx context.Context, email string) (models.RecoverInfo, error)
	SaveVerification(ctx context.Context, info models.VerInfo) error
}

type Actions interface {
	GetUserInfoByID(ctx context.Context, uuid string) (models.UserInfo, error)
	GetUserInfoByEmail(ctx context.Context, email string) (models.UserInfo, error)
	GetUserPassword(ctx context.Context, uuid string) (string, error)
}

type LoginInfo struct {
	UUID     string `json:"uuid"`
	Email    string
	Password string
	Username string
}
