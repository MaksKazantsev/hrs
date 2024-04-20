package service

import (
	"context"
	"fmt"
	"github.com/alserov/hrs/auth/internal/db"
	"github.com/alserov/hrs/auth/internal/models"
	"github.com/alserov/hrs/auth/internal/utils"
	"github.com/google/uuid"
)

type Service interface {
	SignUp(ctx context.Context, req models.RegReq) (models.RegRes, error)
	SignIn(ctx context.Context, req models.LoginReq) (string, error)
}

func NewService(repo db.Repository) Service {
	return &service{
		repo: repo,
	}
}

type service struct {
	repo db.Repository
}

func (s *service) SignUp(ctx context.Context, req models.RegReq) (models.RegRes, error) {
	// generating uuid
	req.UUID = uuid.New().String()

	// generating password hash
	hash, err := utils.GenerateHash(req.Password)
	if err != nil {
		return models.RegRes{}, fmt.Errorf("failed to hash password: %w", err)
	}
	req.Password = hash

	// generating token
	token, err := utils.NewToken(req.UUID)
	if err != nil {
		return models.RegRes{}, fmt.Errorf("failed to generate token: %w", err)
	}

	if err = s.repo.SignUp(ctx, req); err != nil {
		return models.RegRes{}, fmt.Errorf("repo error: %w", err)
	}
	return models.RegRes{
		UUID:  req.UUID,
		Token: token,
	}, nil
}

func (s *service) SignIn(ctx context.Context, req models.LoginReq) (string, error) {
	info, err := s.repo.SignIn(ctx, req.Email)
	if err != nil {
		return "", fmt.Errorf("repo error: %w", err)
	}

	// comparing password
	if err = utils.CompareHash(info.Password, req.Password); err != nil {
		return "", fmt.Errorf("error: %w", err)
	}

	// generating token
	token, err := utils.NewToken(info.UUID)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}
	return token, nil
}
