package server

import (
	"context"
	"github.com/alserov/hrs/auth/internal/log"
	"github.com/alserov/hrs/auth/internal/service"
	"github.com/alserov/hrs/auth/internal/utils"
	"github.com/alserov/hrs/auth/internal/utils/converter"
	"github.com/alserov/hrs/auth/internal/utils/validator"
	"github.com/alserov/hrs/auth/pkg/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	gen.UnimplementedUserServer

	converter converter.Converter

	validator validator.Validator

	log log.Logger

	service service.Service
}

func RegisterGRPCServer(s *grpc.Server, service service.Service) {
	gen.RegisterUserServer(s, newServer(service))
}

func newServer(s service.Service) gen.UserServer {
	return &server{
		converter: converter.NewConverter(),
		validator: validator.NewValidator(),
		log:       log.GetLogger(),
		service:   s,
	}
}

func (s *server) Register(ctx context.Context, req *gen.RegisterReq) (*gen.RegisterRes, error) {
	if err := s.validator.ValidateRegReq(req); err != nil {
		return nil, err
	}

	data, err := s.service.SignUp(ctx, s.converter.RegReqToService(req))
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return s.converter.RegResToPb(data), nil
}

func (s *server) Login(ctx context.Context, req *gen.LoginReq) (*gen.LoginRes, error) {
	if err := s.validator.ValidateLoginReq(req); err != nil {
		return nil, err
	}

	token, err := s.service.SignIn(ctx, s.converter.LoginReqToService(req))
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return s.converter.LoginResToPb(token), nil
}

func (s *server) Reset(ctx context.Context, req *gen.ResetReq) (*emptypb.Empty, error) {
	if err := s.validator.ValidateResReq(req); err != nil {
		return nil, err
	}
	if err := s.service.ResetPass(ctx, s.converter.ResetReqToService(req)); err != nil {
		return nil, utils.HandleError(err)
	}

	return nil, nil
}
