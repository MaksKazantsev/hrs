package utils

import (
	"errors"
	"github.com/alserov/hrs/auth/internal/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	ErrInternal = iota + 1
	ErrNotFound
	ErrBadRequest
)
const internalError = "unknown internal error"

type Error struct {
	Status  int
	Message string
}

func (e Error) Error() string {
	return e.Message
}

func NewError(st int, message string) error {
	return &Error{
		Status:  st,
		Message: message,
	}
}

func HandleError(err error) error {
	var e *Error
	l := log.GetLogger()

	if !errors.As(err, &e) {
		return status.Error(codes.Internal, internalError)
	}

	switch e.Status {
	case ErrInternal:
		l.Error(e.Message)
		return status.Error(codes.Internal, "internal error")
	case ErrBadRequest:
		l.Error(e.Message)
		return status.Error(codes.InvalidArgument, e.Message)
	case ErrNotFound:
		return status.Error(codes.NotFound, e.Message)
	default:
		return status.Error(codes.Internal, internalError)
	}
}