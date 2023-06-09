package url

import "context"

//UseCase interface for usecase
type UseCase interface {
	Save(ctx context.Context, long, short string) (string, error)
	Get(ctx context.Context, short string) (string, error)
}
