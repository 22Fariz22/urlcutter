package url

import (
	"context"
	"github.com/22Fariz22/urlcutter/pkg/logger"
)

//go:generate mockgen -source usecase.go -destination mock/usecase.go -package mock

// UseCase interface for usecase
type UseCase interface {
	Save(ctx context.Context, l logger.Interface, long, short string) (string, error)
	Get(ctx context.Context, l logger.Interface, short string) (string, error)
}
