package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gleblagov/electronics-shop/data"
)

type ProductStorage interface {
	GetById(ctx context.Context, id int) (data.Product, error)
	New(ctx context.Context, Product data.Product) (data.Product, error)
	Delete(ctx context.Context, id int) error
}

func HandleProductGetById(ctx context.Context, ps ProductStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idRaw := r.PathValue("id")
		id, err := strconv.Atoi(idRaw)
		// TODO: validation
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		product, err := ps.GetById(ctx, id)
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusInternalServerError)
			return
		}
		enc := json.NewEncoder(w)
		enc.Encode(product)
		return
	}
}

func HandleProductNew(ctx context.Context, ps ProductStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product data.Product
		// TODO: validation
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		newProduct, err := ps.New(ctx, product)
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusInternalServerError)
			return
		}
		enc := json.NewEncoder(w)
		enc.Encode(newProduct)
		return
	}
}

func HandleProductDelete(ctx context.Context, ps ProductStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idRaw := r.PathValue("id")
		id, err := strconv.Atoi(idRaw)
		// TODO: validation
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusBadRequest)
			return
		}

		err = ps.Delete(ctx, id)
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		return
	}
}
