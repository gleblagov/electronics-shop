package data

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
		return nil, err
	}
	return &PostgresStorage{
		Users:    userStorage,
		Products: productStorage,
		Carts:    cartStorage,
	}, nil
}
