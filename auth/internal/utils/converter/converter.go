package converter

import (
	"github.com/alserov/hrs/auth/internal/models"
	"github.com/alserov/hrs/auth/pkg/proto/gen"
)

type Converter interface {
	ToPb
	ToService
}

type ToPb interface {
	RegResToPb(req models.RegRes) *gen.RegisterRes
	LoginResToPb(token string) *gen.LoginRes
}

type ToService interface {
	RegReqToService(req *gen.RegisterReq) models.RegReq
	LoginReqToService(req *gen.LoginReq) models.LoginReq
	ResetReqToService(req *gen.ResetReq) models.ResetReq
}

func NewConverter() Converter {
	return &converter{}
}

type converter struct {
}

func (c *converter) ResetReqToService(req *gen.ResetReq) models.ResetReq {
	return models.ResetReq{
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
		Token:       req.Token,
	}
}

func (c *converter) LoginResToPb(token string) *gen.LoginRes {
	return &gen.LoginRes{
		Token: token,
	}
}

func (c *converter) RegResToPb(req models.RegRes) *gen.RegisterRes {
	return &gen.RegisterRes{
		UUID:  req.UUID,
		Token: req.Token,
	}
}

func (c *converter) LoginReqToService(req *gen.LoginReq) models.LoginReq {
	return models.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (c *converter) RegReqToService(req *gen.RegisterReq) models.RegReq {
	return models.RegReq{
		Email:    req.Email,
		Password: req.Password,
		UserName: req.Username,
	}
}
