package data

type PostgresStorage struct {
	Users *userStoragePostgres
}

func NewPostgresStorage() (*PostgresStorage, error) {
	userStorage, err := newUserStoragePostgres()
	if err != nil {
		return nil, err
	}
	return &PostgresStorage{
		Users: userStorage,
	}, nil
}
