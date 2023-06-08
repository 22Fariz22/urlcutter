package url

import "context"

type UseCase interface {
	Save(ctx context.Context)
	Get(ctx context.Context)
}
