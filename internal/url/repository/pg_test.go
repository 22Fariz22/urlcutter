package repository

import (
	"context"
	"github.com/22Fariz22/urlcutter/pkg/logger"
	"github.com/22Fariz22/urlcutter/pkg/postgres"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_pgRepository_Save(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	pg, err := postgres.New("")
	require.NoError(t, err)

	pgRepo := NewPGRepository(pg)
	defer pgRepo.Close()

	columns := []string{"short"}
	rows := sqlmock.NewRows(columns).AddRow(
		"mock_short_url",
	)
	l := logger.New("debug")
	mock.ExpectQuery(saveURLQuery).WithArgs("mock_long_url", "mock_short_url").WillReturnRows(rows)
	mock.ExpectQuery(selectExistURLQuery).WithArgs("mock_short_url").WillReturnRows(rows)

	value, err := pgRepo.Save(context.Background(), l, "mock_long_url", "mock_short_url")

	require.NoError(t, err)
	require.NotNil(t, value)

}

func Test_pgRepository_Get(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	pg, err := postgres.New("")
	require.NoError(t, err)

	pgRepo := NewPGRepository(pg)
	defer pgRepo.Close()

	columns := []string{"long"}
	rows := sqlmock.NewRows(columns).AddRow(
		"mock_long_url",
	)

	mock.ExpectQuery(getURLQuery).WithArgs("mock_short_url").WillReturnRows(rows)

	long, err := pgRepo.Get(context.Background(), nil, "mock_short_url")

	require.NotNil(t, err)
	require.NotNil(t, long)

}
