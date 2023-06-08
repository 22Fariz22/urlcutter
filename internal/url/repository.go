package url

import "context"

type Repo interface {
	Save(ctx context.Context)
	Get(ctx context.Context)
}
