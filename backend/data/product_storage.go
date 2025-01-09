package data

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type productStoragePostgres struct {
	pool *pgxpool.Pool
}

func newProductStoragePostgres() (*productStoragePostgres, error) {
	slog.Info("Initializing new ProductStoragePostgres...")
	poolconn, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		slog.Error("Failed to connect to Postgres", "op", "newProductStoragePostgres()", "err", err.Error())
		return nil, err
	}
	slog.Info("Initialized productStoragePostgres successfuly")
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
	err := row.Scan(&product.Id, &product.Title, &product.Price, &product.Quantity, &product.Category, &product.Rating)
	// TODO: refactor
	if err != nil {
		if err == pgx.ErrNoRows {
			return Product{}, fmt.Errorf("product with id %d does not exist: %v", id, err)
		}
		slog.Error("Failed to execute SQL query (get product by id)", "op", "productStoragePostgres.GetById()", "err", err.Error())
		return Product{}, fmt.Errorf("failed to execute query: %v", err)
	}
	return product, nil
}

func (ps productStoragePostgres) New(ctx context.Context, product Product) (Product, error) {
	q := `
        INSERT INTO products (price, title, quantity, category, rating)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, title, price, quantity, category, rating
    `
	var createdProduct Product
	err := ps.pool.QueryRow(ctx, q, product.Price, product.Title, product.Quantity, product.Category, product.Rating).
		Scan(&createdProduct.Id, &createdProduct.Title, &createdProduct.Price, &createdProduct.Quantity, &createdProduct.Category, &createdProduct.Rating)
		// TODO: refactor
	if err != nil {
		slog.Error("Failed to execute SQL query (insert new product)", "op", "productStoragePostgres.New()", "err", err.Error())
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
		slog.Error("Failed to execute SQL query (delete product)", "op", "productStoragePostgres.New()", "err", err.Error())
		return fmt.Errorf("failed to delete product with id %d: %w", id, err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("product with id %d does not exist", id)
	}

	return nil
}
