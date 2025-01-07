package data

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type cartStoragePostgres struct {
	pool *pgxpool.Pool
}

func newCartStoragePostgres() (*cartStoragePostgres, error) {
	poolconn, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		return nil, err
	}
	return &cartStoragePostgres{
		pool: poolconn,
	}, nil
}

func (cs cartStoragePostgres) GetById(ctx context.Context, id int) (Cart, error) {
	q := `
    SELECT *
    FROM carts
    WHERE id = $1
    `
	var cart Cart
	row := cs.pool.QueryRow(ctx, q, id)
	err := row.Scan(&cart.Id, &cart.UserId, &cart.CreatedAt, &cart.UpdatedAt, cart.Status)
	// TODO: refactor
	if err != nil {
		if err == pgx.ErrNoRows {
			return Cart{}, fmt.Errorf("cart with id %d does not exist: %v", id, err)
		}
		return Cart{}, fmt.Errorf("failed to execute query: %v", err)
	}
	return cart, nil
}

func (cs cartStoragePostgres) New(ctx context.Context, userId int) (Cart, error) {
	q := `
        INSERT INTO carts (user_id)
        VALUES ($1)
        RETURNING id, user_id, created_at, updated_at, status
    `
	var createdCart Cart
	err := cs.pool.QueryRow(ctx, q, userId).
		Scan(&createdCart.Id, &createdCart.UserId, &createdCart.CreatedAt, &createdCart.UpdatedAt, &createdCart.Status)
		// TODO: refactor
	if err != nil {
		return Cart{}, fmt.Errorf("failed to execute query: %v", err)
	}
	return createdCart, nil
}

func (cs cartStoragePostgres) ChangeStatus(ctx context.Context, id int, status string) error {
	allowedStatuses := []string{"created", "closed", "purchased"}
	allowed := false
	for _, s := range allowedStatuses {
		if status == s {
			allowed = true
			break
		}
	}
	if !allowed {
		return fmt.Errorf("invalid status: %s", status)
	}
	q := `
        UPDATE carts
        SET status = $1
		WHERE id = $2
        RETURNING id, user_id, created_at, updated_at, status
    `
	tag, err := cs.pool.Exec(ctx, q, status, id)
	if err != nil {
		return fmt.Errorf("failed update cart status with cart id %d: %w", id, err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("cart with id %d does not exist", id)
	}
	return nil
}
