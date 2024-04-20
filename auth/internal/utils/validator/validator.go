package validator

import (
	"github.com/alserov/hrs/auth/pkg/proto/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
)

const (
	ERR_INVALID_PASSWORD = "invalid password"
	ERR_INVALID_USERNAME = "invalid username"
	ERR_INVALID_EMAIL    = "invalid email"
)

type Validator interface {
	ValidateRegReq(req *gen.RegisterReq) error
	ValidateLoginReq(req *gen.LoginReq) error
	ValidateResReq(req *gen.ResetReq) error
}

func NewValidator() Validator {
	return &validator{
		regExpEmail: regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`),
	}
}

type validator struct {
	regExpEmail *regexp.Regexp
}

func (v validator) ValidateLoginReq(req *gen.LoginReq) error {
	if ok := v.regExpEmail.MatchString(req.Email); !ok {
		return status.Error(codes.InvalidArgument, ERR_INVALID_EMAIL)
	}
	if err := ValidatePassword(req.Password); err != nil {
		return err
	}
	return nil
}

func (v validator) ValidateRegReq(req *gen.RegisterReq) error {
	if len(req.Username) > 20 || len(req.Username) < 2 {
		return status.Error(codes.InvalidArgument, ERR_INVALID_USERNAME)
	}

	if ok := v.regExpEmail.MatchString(req.Email); !ok {
		return status.Error(codes.InvalidArgument, ERR_INVALID_EMAIL)
	}

	if err := ValidatePassword(req.Password); err != nil {
		return err
	}
	return nil
}

func (v validator) ValidateResReq(req *gen.ResetReq) error {
	if err := ValidatePassword(req.OldPassword); err != nil {
		return status.Error(codes.InvalidArgument, ERR_INVALID_PASSWORD)
	}
	if err := ValidatePassword(req.NewPassword); err != nil {
		return status.Error(codes.InvalidArgument, ERR_INVALID_PASSWORD)
	}
	return nil
}

func ValidatePassword(pass string) error {
	if len(pass) < 7 || len(pass) > 40 {
		return status.Error(codes.InvalidArgument, ERR_INVALID_PASSWORD)
	}
	return nil
}
