package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/22Fariz22/urlcutter/pkg/grpcerrors"
	"github.com/22Fariz22/urlcutter/pkg/postgres"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

type PGRepository struct {
	*postgres.Postgres
}

func NewPGRepository(db *postgres.Postgres) *PGRepository {
	return &PGRepository{db}
}

func (p *PGRepository) Save(ctx context.Context, long, short string) (string, error) {
	fmt.Println("here PG repo Save()")
	var alreadyExistValue string

	//вставляем урл, если такой уже существует, то вернем ошибку.Если новый урл,то вернем шортурл
	err := p.Pool.QueryRow(ctx,
		`insert into urls (long,short) values ($1,$2) returning short`, long, short).Scan(&alreadyExistValue)

	var pgErr *pgconn.PgError
	if err != nil {
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			//делаем запрос чтобы вернуть существующий шортурл
			_ = p.Pool.QueryRow(context.Background(),
				`SELECT short FROM urls where long = $1;`, long).Scan(&alreadyExistValue)
			fmt.Println("вернем уже сущемтвующий:", alreadyExistValue)
			return alreadyExistValue, grpcerrors.ErrURLExists
		}
	}

	fmt.Println("alreadyExistValue", alreadyExistValue)
	return alreadyExistValue, nil
}

func (p *PGRepository) Get(ctx context.Context, short string) (string, error) {
	fmt.Println("here PG repo Get()")

	return "", nil
}
