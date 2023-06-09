package repository

import (
	"context"
	"errors"
	"fmt"
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

//Save url to db
func (p *pgRepository) Save(ctx context.Context, long, short string) (string, error) {
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

//Get url from db
func (p *pgRepository) Get(ctx context.Context, short string) (string, error) {
	fmt.Println("here PG repo Get()")

	var existLong string
	err := p.Pool.QueryRow(ctx, "select long from urls where short=$1", short).Scan(&existLong)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", grpcerrors.ErrDoesNotExist
		}
		return "", grpcerrors.ErrPG
	}

	fmt.Println("existLong:", existLong)

	return existLong, nil
}
