package repository

import (
	"context"
	"fmt"
	"github.com/22Fariz22/urlcutter/pkg/postgres"
)

type PGRepository struct {
	*postgres.Postgres
}

func NewPGRepository(db *postgres.Postgres) *PGRepository {
	return &PGRepository{db}
}

func (p *PGRepository) Save(ctx context.Context) {
	fmt.Println("here PG repo Save()")
}

func (p *PGRepository) Get(ctx context.Context) {
	fmt.Println("here PG repo Get()")

}
