package url

import (
	"context"
	"github.com/22Fariz22/urlcutter/pkg/logger"
)

//go:generate mockgen -source repository.go -destination mock/repository.go -package mock

// Repo interface for storages
type Repo interface {
	Save(ctx context.Context, l logger.Interface, long, short string) (string, error)
	Get(ctx context.Context, l logger.Interface, short string) (string, error)
}
