package repository

import (
	"context"
	"errors"
	"github.com/22Fariz22/urlcutter/pkg/logger"

	"github.com/22Fariz22/urlcutter/pkg/grpcerrors"
	"github.com/22Fariz22/urlcutter/pkg/postgres"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type pgRepository struct {
	*postgres.Postgres
}

// NewPGRepository create postgres storage
func NewPGRepository(db *postgres.Postgres) *pgRepository {
	return &pgRepository{db}
}

// Save url to db
func (p *pgRepository) Save(ctx context.Context, l logger.Interface, long, short string) (string, error) {
	var alreadyExistValue string

	//вставляем урл, если такой уже существует, то вернем ошибку.Если новый урл,то вернем шортурл
	err := p.Pool.QueryRow(ctx,
		saveURLQuery, long, short).Scan(&alreadyExistValue)

	var pgErr *pgconn.PgError
	if err != nil {
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			//делаем запрос чтобы вернуть существующий шортурл
			_ = p.Pool.QueryRow(context.Background(),
				selectExistURLQuery, long).Scan(&alreadyExistValue)
			return alreadyExistValue, grpcerrors.ErrURLExists
		}
	}

	return alreadyExistValue, nil
}

// Get url from db
func (p *pgRepository) Get(ctx context.Context, l logger.Interface, short string) (string, error) {
	var existLong string
	err := p.Pool.QueryRow(ctx, getURLQuery, short).Scan(&existLong)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", grpcerrors.ErrDoesNotExist
		}
		return "", grpcerrors.ErrPG
	}

	return existLong, nil
}
