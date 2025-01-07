package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gleblagov/electronics-shop/data"
)

type CartStorage interface {
	GetById(ctx context.Context, id int) (data.Cart, error)
	New(ctx context.Context, userId int) (data.Cart, error)
	ChangeStatus(ctx context.Context, id int, status string) (data.Cart, error)
	AddProductToCart(ctx context.Context, cartId int, productId int, quantity int) (data.CartItem, error)
}

func HandleCartGetById(ctx context.Context, cs CartStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idRaw := r.PathValue("id")
		id, err := strconv.Atoi(idRaw)
		// TODO: validation
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		cart, err := cs.GetById(ctx, id)
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusInternalServerError)
			return
		}
		enc := json.NewEncoder(w)
		enc.Encode(cart)
		return
	}
}

func HandleCartNew(ctx context.Context, cs CartStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cartReq := struct {
			UserId int `json:"user_id"`
		}{}
		// TODO: validation
		if err := json.NewDecoder(r.Body).Decode(&cartReq); err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		newCart, err := cs.New(ctx, cartReq.UserId)
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusInternalServerError)
			return
		}
		enc := json.NewEncoder(w)
		enc.Encode(newCart)
		return
	}
}

func HandleCartChangeStatus(ctx context.Context, cs CartStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idRaw := r.PathValue("id")
		id, err := strconv.Atoi(idRaw)
		// TODO: validation
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusBadRequest)
			return
		}
		cartReq := struct {
			Status string `json:"status"`
		}{}
		if err := json.NewDecoder(r.Body).Decode(&cartReq); err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		updatedCart, err := cs.ChangeStatus(ctx, id, cartReq.Status)
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusInternalServerError)
			return
		}
		enc := json.NewEncoder(w)
		enc.Encode(updatedCart)
		return
	}
}

func HandleCartAddProduct(ctx context.Context, cs CartStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idRaw := r.PathValue("id")
		id, err := strconv.Atoi(idRaw)
		// TODO: validation
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusBadRequest)
			return
		}
		prodReq := struct {
			ProductId int `json:"product_id"`
			Quantity  int `json:"quantity"`
		}{}
		if err := json.NewDecoder(r.Body).Decode(&prodReq); err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		addedItem, err := cs.AddProductToCart(ctx, id, prodReq.ProductId, prodReq.Quantity)
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusInternalServerError)
			return
		}
		enc := json.NewEncoder(w)
		enc.Encode(addedItem)
		return
	}
}
