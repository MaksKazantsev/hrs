package service

import (
	"context"
	"github.com/alserov/hrs/auth/internal/db"
	"github.com/alserov/hrs/auth/internal/models"
	"github.com/alserov/hrs/auth/internal/utils"
)

type Service interface {
	SignUp(ctx context.Context, req models.RegReq) (models.RegRes, error)
	SignIn(ctx context.Context, req models.LoginReq) (string, error)
	ResetPass(ctx context.Context, req models.ResetReq) error
	RecoverPass(ctx context.Context, req models.RecoverReq) (string, error)
	Verify(ctx context.Context, req models.VerifyReq) error
}

func NewService(repo db.Repository) Service {
	return &service{
		repo:   repo,
		sender: utils.NewCodeSender(),
	}
}

type service struct {
	repo   db.Repository
	sender utils.CodeSender
}
