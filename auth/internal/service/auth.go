package service

import (
	"context"
	"fmt"
	"github.com/alserov/hrs/auth/internal/log"
	"github.com/alserov/hrs/auth/internal/models"
	"github.com/alserov/hrs/auth/internal/utils"
	"github.com/google/uuid"
	"math/rand"
	"strconv"
)

const (
	verification = "verif"
	recoverPass  = "recover"
)

func (s *service) SignUp(ctx context.Context, req models.RegReq) (models.RegRes, error) {
	// logging
	log.GetLogger(ctx).Debug("usecase layer success ✔")

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

	// calling repo method
	if err = s.repo.SignUp(ctx, req); err != nil {
		return models.RegRes{}, fmt.Errorf("repo error: %w", err)
	}

	// send verification code
	code := strconv.Itoa(rand.Intn(8999) + 1000)
	if err = s.sender.SendCode(code, req.Email); err != nil {
		return models.RegRes{}, fmt.Errorf("failed to send code")
	}
	// calling repo method
	info := models.VerInfo{
		Email:      req.Email,
		Code:       code,
		IsVerified: false,
	}
	if err = s.repo.SaveVerification(ctx, info); err != nil {
		return models.RegRes{}, fmt.Errorf("repo error: %w", err)
	}

	return models.RegRes{
		UUID:  req.UUID,
		Token: token,
	}, nil
}

func (s *service) SignIn(ctx context.Context, req models.LoginReq) (string, error) {
	// logging
	log.GetLogger(ctx).Debug("usecase layer success ✔")

	// calling repo method
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

func (s *service) ResetPass(ctx context.Context, req models.ResetReq) error {
	// logging
	log.GetLogger(ctx).Debug("usecase layer success ✔")

	// Parse token
	userID, err := utils.ParseToken(req.Token)
	if err != nil {
		return fmt.Errorf("failed to parse token: %w", err)
	}

	// getting info about user
	pass, err := s.repo.GetUserPassword(ctx, userID)
	if err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	// comparing passwords
	if err = utils.CompareHash(pass, req.OldPassword); err != nil {
		return fmt.Errorf("wrong password: %w", err)
	}
	// hashing new password
	hash, err := utils.GenerateHash(req.NewPassword)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// resetting password
	err = s.repo.ResetPass(ctx, userID, hash)
	if err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	return nil
}

func (s *service) RecoverPass(ctx context.Context, req models.RecoverReq) (string, error) {
	// logging
	log.GetLogger(ctx).Debug("usecase layer success ✔")

	// getting user id and check if verified
	user, err := s.repo.GetUserInfoByEmail(ctx, req.Email)
	if err != nil {
		return "", fmt.Errorf("repo error: %w", err)
	}

	if !user.IsVerified {
		return "", fmt.Errorf("verify your account")
	}

	// hashing password
	pass, err := utils.GenerateHash(req.NewPassword)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	req.NewPassword = pass

	// send recover code
	code := strconv.Itoa(rand.Intn(8999) + 1000)
	if err = s.sender.SendCode(code, req.Email); err != nil {
		return "", fmt.Errorf("failed to send code")
	}
	req.Code = code

	// calling repo method
	if err = s.repo.RecoverPass(ctx, req); err != nil {
		return "", fmt.Errorf("repo error: %w", err)
	}

	// generating token
	token, err := utils.NewToken(user.UUID)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}
	return token, nil
}

func (s *service) Verify(ctx context.Context, req models.VerifyReq) error {
	switch req.Typo {
	case verification:
		// logging
		log.GetLogger(ctx).Debug("usecase layer success ✔")

		// calling repo method
		info, err := s.repo.GetVerification(ctx, req.Email)
		if err != nil {
			return fmt.Errorf("repo error: %w", err)
		}

		// calling verification repo method
		if err = s.repo.Verificate(ctx, info.Code, req.Email); err != nil {
			return fmt.Errorf("repo error: %v", err)
		}
	case recoverPass:
		// logging
		log.GetLogger(ctx).Debug("usecase layer success ✔")

		// calling repo method
		info, err := s.repo.GetRecover(ctx, req.Email)
		if err != nil {
			return fmt.Errorf("repo error: %w", err)
		}

		// calling verification repo method
		if err = s.repo.VerificateRecover(ctx, info.Code, req.Email, info.Password); err != nil {
			return fmt.Errorf("repo error: %v", err)
		}
	}
	return nil
}
