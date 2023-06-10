package url

import "context"

//go:generate mockgen -source repository.go -destination mock/repository.go -package mock

// Repo interface for storages
type Repo interface {
	Save(ctx context.Context, long, short string) (string, error)
	Get(ctx context.Context, short string) (string, error)
}
