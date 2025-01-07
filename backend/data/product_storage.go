package data

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type productStoragePostgres struct {
	pool *pgxpool.Pool
}

func newProductStoragePostgres() (*productStoragePostgres, error) {
	poolconn, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		return nil, err
	}
	return &productStoragePostgres{
		pool: poolconn,
	}, nil
}

func (ps productStoragePostgres) GetById(ctx context.Context, id int) (Product, error) {
	q := `
    SELECT *
    FROM products
    WHERE id = $1
    `
	var product Product
	row := ps.pool.QueryRow(ctx, q, id)
	err := row.Scan(&product.Id, &product.Price, &product.Quantity, &product.Category, &product.Rating)
	// TODO: refactor
	if err != nil {
		if err == pgx.ErrNoRows {
			return Product{}, fmt.Errorf("product with id %d does not exist: %v", id, err)
		}
		return Product{}, fmt.Errorf("failed to execute query: %v", err)
	}
	return product, nil
}

func (ps productStoragePostgres) New(ctx context.Context, product Product) (Product, error) {
	q := `
        INSERT INTO products (price, title, quantity, category)
        VALUES ($1, $2, $3)
        RETURNING id, title, price, quantity, category, rating
    `
	var createdProduct Product
	err := ps.pool.QueryRow(ctx, q, product.Price, product.Title, product.Quantity, product.Category).
		Scan(&createdProduct.Id, &createdProduct.Title, &createdProduct.Price, &createdProduct.Quantity, &createdProduct.Category, &createdProduct.Rating)
		// TODO: refactor
	if err != nil {
		return Product{}, fmt.Errorf("failed to execute query: %v", err)
	}
	return createdProduct, nil
}

func (ps productStoragePostgres) Delete(ctx context.Context, id int) error {
	q := `
        DELETE FROM products
        WHERE id = $1
    `
	tag, err := ps.pool.Exec(ctx, q, id)
	if err != nil {
		return fmt.Errorf("failed to delete product with id %d: %w", id, err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("product with id %d does not exist", id)
	}

	return nil
}
