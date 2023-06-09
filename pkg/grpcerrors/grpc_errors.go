package grpcerrors

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
)

var (
	//ErrNotFound не найдено
	ErrNotFound = errors.New("Not found")
	//ErrURLExists URL exists
	ErrURLExists = errors.New("URL already exists")
	//ErrDoesNotExist url not found
	ErrDoesNotExist = errors.New("this URL does not exist")

	//ErrPG error from select db
	ErrPG = errors.New("error from select PG")
)

// ParseGRPCErrStatusCode Parse error and get code
func ParseGRPCErrStatusCode(err error) codes.Code {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return codes.NotFound
	case errors.Is(err, context.Canceled):
		return codes.Canceled
	case errors.Is(err, context.DeadlineExceeded):
		return codes.DeadlineExceeded
	case errors.Is(err, ErrURLExists):
		return codes.AlreadyExists

	case errors.Is(err, ErrDoesNotExist):
		return codes.AlreadyExists
	}
	return codes.Internal
}
