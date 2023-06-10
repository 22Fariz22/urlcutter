package url

import "context"

//go:generate mockgen -source usecase.go -destination mock/usecase.go -package mock

// UseCase interface for usecase
type UseCase interface {
	Save(ctx context.Context, long, short string) (string, error)
	Get(ctx context.Context, short string) (string, error)
}
