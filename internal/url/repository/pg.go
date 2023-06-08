package repository

import "github.com/22Fariz22/urlcutter/pkg/postgres"

type PGRepository struct {
	*postgres.Postgres
}

func NewPGRepository(db *postgres.Postgres) *PGRepository {
	return &PGRepository{db}
}

func (p *PGRepository) Save() {

}

func (p *PGRepository) Get() {

}
