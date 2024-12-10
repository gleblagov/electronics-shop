package data

type PostgresStorage struct {
	us *userStoragePostgres
}

func NewPostgresStorage() (*PostgresStorage, error) {
	userStorage, err := newUserStoragePostgres()
	if err != nil {
		return nil, err
	}
	return &PostgresStorage{
		us: userStorage,
	}, nil
}
