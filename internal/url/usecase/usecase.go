package usecase

import (
	"context"
	"fmt"
	"github.com/22Fariz22/urlcutter/internal/url"
)

type UseCase struct {
	repo url.Repo
}

func NewUseCase(repo url.Repo) *UseCase {
	return &UseCase{repo: repo}
}

func (u *UseCase) Save(ctx context.Context, long, short string) (string, error) {
	fmt.Println("here UC Save()")
	return u.repo.Save(ctx, long, short)
}

func (u UseCase) Get(ctx context.Context, short string) (string, error) {
	fmt.Println("here UC Get()")
	return u.repo.Get(ctx, short)
}
