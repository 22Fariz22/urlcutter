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

func (u *UseCase) Save(ctx context.Context) {
	fmt.Println("here UC Save()")
	u.repo.Save(ctx)
}

func (u UseCase) Get(ctx context.Context) {
	fmt.Println("here UC Get()")
	u.repo.Get(ctx)
}
