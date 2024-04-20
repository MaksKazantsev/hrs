package service

import (
	"context"
	"github.com/alserov/hrs/auth/internal/db"
	"github.com/alserov/hrs/auth/internal/models"
)

type Service interface {
	SignUp(ctx context.Context, req models.RegReq) (models.RegRes, error)
	SignIn(ctx context.Context, req models.LoginReq) (string, error)
	ResetPass(ctx context.Context, req models.ResetReq) error
}

func NewService(repo db.Repository) Service {
	return &service{
		repo: repo,
	}
}

type service struct {
	repo db.Repository
}
