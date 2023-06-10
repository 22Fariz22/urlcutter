package usecase

import (
	"context"

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
func (u *useCase) Save(ctx context.Context, long, short string) (string, error) {
	return u.repo.Save(ctx, long, short)
}

// Get delivery url to usecase method
func (u useCase) Get(ctx context.Context, short string) (string, error) {
	return u.repo.Get(ctx, short)
}
