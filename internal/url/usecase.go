package url

import "context"

type UseCase interface {
	Save(ctx context.Context, long, short string) (string, error)
	Get(ctx context.Context, short string) (string, error)
}
