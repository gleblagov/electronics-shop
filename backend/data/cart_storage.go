package data

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductStorage interface {
	GetById(ctx context.Context, id int) (Product, error)
}

type cartStoragePostgres struct {
	pool *pgxpool.Pool
	ps   ProductStorage
}

func newCartStoragePostgres(ps ProductStorage) (*cartStoragePostgres, error) {
	poolconn, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		return nil, err
	}
	return &cartStoragePostgres{
		pool: poolconn,
		ps:   ps,
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
	err := row.Scan(&cart.Id, &cart.UserId, &cart.CreatedAt, &cart.UpdatedAt, &cart.Status)
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
	var userExists bool
	checkUserQuery := `SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)`
	err := cs.pool.QueryRow(ctx, checkUserQuery, userId).Scan(&userExists)
	if err != nil {
		return Cart{}, fmt.Errorf("failed to execute query: %v", err)
	}
	if !userExists {
		return Cart{}, fmt.Errorf("user with id %d does not exist", userId)
	}
	q := `
        INSERT INTO carts (user_id)
        VALUES ($1)
        RETURNING id, user_id, created_at, updated_at, status
    `
	var createdCart Cart
	err = cs.pool.QueryRow(ctx, q, userId).
		Scan(&createdCart.Id, &createdCart.UserId, &createdCart.CreatedAt, &createdCart.UpdatedAt, &createdCart.Status)
		// TODO: refactor
	if err != nil {
		return Cart{}, fmt.Errorf("failed to execute query: %v", err)
	}
	return createdCart, nil
}

func (cs cartStoragePostgres) ChangeStatus(ctx context.Context, id int, status string) (Cart, error) {
	allowedStatuses := []string{"created", "closed", "purchased"}
	allowed := false
	for _, s := range allowedStatuses {
		if status == s {
			allowed = true
			break
		}
	}
	if !allowed {
		return Cart{}, fmt.Errorf("invalid status: %s", status)
	}
	q := `
        UPDATE carts
        SET status = $1, updated_at = CURRENT_TIMESTAMP
		WHERE id = $2
        RETURNING id, user_id, created_at, updated_at, status
    `
	var updatedCart Cart
	err := cs.pool.QueryRow(ctx, q, status, id).Scan(&updatedCart.Id, &updatedCart.UserId, &updatedCart.CreatedAt, &updatedCart.UpdatedAt, &updatedCart.Status)
	if err != nil {
		return Cart{}, fmt.Errorf("failed update cart status with cart id %d: %w", id, err)
	}

	return updatedCart, nil
}

func (cs cartStoragePostgres) AddProductToCart(ctx context.Context, cartId int, productId int, quantity int) (CartItem, error) {
	product, err := cs.ps.GetById(ctx, productId)
	if err != nil {
		return CartItem{}, fmt.Errorf("failed to find product with id %d: %w", productId, err)
	}
	_, err = cs.GetById(ctx, cartId)
	if err != nil {
		return CartItem{}, fmt.Errorf("failed to find cart with id %d: %w", cartId, err)
	}
	q := `
        INSERT INTO cart_items (cart_id, product_id, quantity, price)
        VALUES ($1, $2, $3, $4)
        RETURNING id, cart_id, product_id, quantity, price, total_cost
    `
	var addedItem CartItem
	err = cs.pool.QueryRow(ctx, q, cartId, productId, quantity, product.Price).
		Scan(&addedItem.Id, &addedItem.CartId, &addedItem.ProductId, &addedItem.Quantity, &addedItem.Price, &addedItem.TotalCost)
	if err != nil {
		return CartItem{}, fmt.Errorf("failed to execute query: %v", err)
	}
	return addedItem, nil
}
