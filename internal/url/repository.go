package url

import "context"

// Repo interface for storages
type Repo interface {
	Save(ctx context.Context, long, short string) (string, error)
	Get(ctx context.Context, short string) (string, error)
}
