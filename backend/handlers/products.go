package handlers

import (
	"context"

	"github.com/gleblagov/electronics-shop/data"
)

type ProductStorage interface {
	GetById(ctx context.Context, id int) (data.Product, error)
	New(ctx context.Context, Product data.Product) (data.Product, error)
	Delete(ctx context.Context, id int) error
}
