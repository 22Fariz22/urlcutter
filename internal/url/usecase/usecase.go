package usecase

import (
	"context"
	"github.com/22Fariz22/urlcutter/pkg/logger"

	"github.com/22Fariz22/urlcutter/internal/url"
)

type useCase struct {
	repo url.Repo
}

// NewUseCase create usecase
func NewUseCase(repo url.Repo) *useCase {
	return &useCase{repo: repo}
}

// Save delivery url to usecase method
func (u *useCase) Save(ctx context.Context, l logger.Interface, long, short string) (string, error) {
	return u.repo.Save(ctx, l, long, short)
}

// Get delivery url to usecase method
func (u useCase) Get(ctx context.Context, l logger.Interface, short string) (string, error) {
	return u.repo.Get(ctx, l, short)
}
