package data

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userStoragePostgres struct {
	pool *pgxpool.Pool
}

func newUserStoragePostgres() (*userStoragePostgres, error) {
	poolconn, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		return nil, err
	}
	return &userStoragePostgres{
		pool: poolconn,
	}, nil
}
