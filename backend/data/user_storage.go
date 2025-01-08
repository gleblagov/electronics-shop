package data

import (
	"context"
	"fmt"
	"os"

	"github.com/gleblagov/electronics-shop/utils"
	"github.com/jackc/pgx/v5"
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

func (us userStoragePostgres) GetById(ctx context.Context, id int) (UserPublic, error) {
	q := `
    SELECT id, email, role
    FROM users
    WHERE id = $1
    `
	var user UserPublic
	row := us.pool.QueryRow(ctx, q, id)
	err := row.Scan(&user.Id, &user.Email, &user.Role)
	// TODO: refactor
	if err != nil {
		if err == pgx.ErrNoRows {
			return UserPublic{}, fmt.Errorf("user with id %d does not exist: %v", id, err)
		}
		return UserPublic{}, fmt.Errorf("failed to execute query: %v", err)
	}
	return user, nil
}

func (us userStoragePostgres) New(ctx context.Context, user User) (UserPublic, error) {
	hashedPass, err := utils.HashPass(user.Password)
	if err != nil {
		return UserPublic{}, fmt.Errorf("failed to hash pass: %v", err)
	}
	q := `
        INSERT INTO users (email, password, role)
        VALUES ($1, $2, $3)
        RETURNING id, email, role
    `
	var createdUser UserPublic
	err = us.pool.QueryRow(ctx, q, user.Email, hashedPass, user.Role).Scan(&createdUser.Id, &createdUser.Email, &createdUser.Role)
	// TODO: refactor
	if err != nil {
		return UserPublic{}, fmt.Errorf("failed to execute query: %v", err)
	}
	return createdUser, nil
}

func (us userStoragePostgres) Update(ctx context.Context, id int, newBody User) (UserPublic, error) {
	hashedPass, err := utils.HashPass(newBody.Password)
	if err != nil {
		return UserPublic{}, fmt.Errorf("failed to hash pass: %v", err)
	}
	q := `
	UPDATE users
	SET email = $1, password = $2, role = $3
	WHERE id = $4
	RETURNING id, email, role
	`
	var updatedUser UserPublic
	err = us.pool.QueryRow(ctx, q, newBody.Email, hashedPass, newBody.Role, id).Scan(&updatedUser.Id, &updatedUser.Email, &updatedUser.Role)
	if err != nil {
		return UserPublic{}, fmt.Errorf("failed to execute query: %v", err)
	}
	return updatedUser, nil
}

func (us userStoragePostgres) Delete(ctx context.Context, id int) error {
	q := `
        DELETE FROM users
        WHERE id = $1
    `
	tag, err := us.pool.Exec(ctx, q, id)
	if err != nil {
		return fmt.Errorf("failed to delete user with id %d: %w", id, err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("user with id %d does not exist", id)
	}

	return nil
}
