package grpcerrors

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"net/http"
)

var (
	//ErrNotFound не найдено
	ErrNotFound = errors.New("Not found")
	//ErrURLExists URL exists
	ErrURLExists = errors.New("URL already exists")
	//ErrBadShortURL short url incorrect, less then 10 symbols
	ErrBadShortURL = errors.New("short url incorrect, less then 10 symbols")
)

// ParseGRPCErrStatusCode Parse error and get code
func ParseGRPCErrStatusCode(err error) codes.Code {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return codes.NotFound
	case errors.Is(err, ErrBadShortURL):
		return codes.InvalidArgument

	case errors.Is(err, context.Canceled):
		return codes.Canceled
	case errors.Is(err, context.DeadlineExceeded):
		return codes.DeadlineExceeded
	case errors.Is(err, ErrURLExists):
		return codes.AlreadyExists
	}
	return codes.Internal
}

// MapGRPCErrCodeToHttpStatus Map GRPC errors codes to http status
func MapGRPCErrCodeToHTTPStatus(code codes.Code) int {
	switch code {
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.AlreadyExists:
		return http.StatusBadRequest
	case codes.NotFound:
		return http.StatusNotFound
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Canceled:
		return http.StatusRequestTimeout
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.InvalidArgument:
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}
