package data

import "log/slog"

type PostgresStorage struct {
	Users    *userStoragePostgres
	Products *productStoragePostgres
	Carts    *cartStoragePostgres
}

func NewPostgresStorage() (*PostgresStorage, error) {
	userStorage, err := newUserStoragePostgres()
	productStorage, err := newProductStoragePostgres()
	cartStorage, err := newCartStoragePostgres(productStorage)
	if err != nil {
		slog.Error("Failed to initialize new Postgres storage", "op", "NewPostgresStorage()", "err", err.Error())
		return nil, err
	}
	return &PostgresStorage{
		Users:    userStorage,
		Products: productStorage,
		Carts:    cartStorage,
	}, nil
}
