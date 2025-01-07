package data

type PostgresStorage struct {
	Users    *userStoragePostgres
	Products *productStoragePostgres
}

func NewPostgresStorage() (*PostgresStorage, error) {
	userStorage, err := newUserStoragePostgres()
	productStorage, err := newProductStoragePostgres()
	if err != nil {
		return nil, err
	}
	return &PostgresStorage{
		Users:    userStorage,
		Products: productStorage,
	}, nil
}
