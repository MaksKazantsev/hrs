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

func NewServer(l log.Logger, s service.Service) *grpc.Server {
	serv := grpc.NewServer()
	gen.RegisterUserServer(serv, &server{converter: converter.NewConverter(), validator: validator.NewValidator(), log: l, service: s})
	return serv
}

func (s *server) Register(ctx context.Context, req *gen.RegisterReq) (*gen.RegisterRes, error) {
	if err := s.validator.ValidateRegReq(req); err != nil {
		return nil, err
	}

	s.log.Debug("received register request ✔")

	data, err := s.service.SignUp(log.WithLogger(ctx, s.log), s.converter.RegReqToService(req))
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return s.converter.RegResToPb(data), nil
}

func (s *server) Login(ctx context.Context, req *gen.LoginReq) (*gen.LoginRes, error) {
	if err := s.validator.ValidateLoginReq(req); err != nil {
		return nil, err
	}

	s.log.Debug("received login request ✔")

	token, err := s.service.SignIn(log.WithLogger(ctx, s.log), s.converter.LoginReqToService(req))
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return s.converter.LoginResToPb(token), nil
}

func (s *server) Reset(ctx context.Context, req *gen.ResetReq) (*emptypb.Empty, error) {
	if err := s.validator.ValidateResReq(req); err != nil {
		return nil, err
	}

	s.log.Debug("received reset request ✔")

	if err := s.service.ResetPass(log.WithLogger(ctx, s.log), s.converter.ResetReqToService(req)); err != nil {
		return nil, utils.HandleError(err)
	}

	return nil, nil
}
