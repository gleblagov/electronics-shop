package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gleblagov/electronics-shop/data"
)

type UserStorage interface {
	GetById(ctx context.Context, id int) (data.UserPublic, error)
	New(ctx context.Context, user data.User) (data.UserPublic, error)
	Delete(ctx context.Context, id int) error
}

func HandleUserGetById(ctx context.Context, us UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idRaw := r.PathValue("id")
		id, err := strconv.Atoi(idRaw)
		// TODO: validation
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		user, err := us.GetById(ctx, id)
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusInternalServerError)
			return
		}
		enc := json.NewEncoder(w)
		enc.Encode(user)
		return
	}
}

func HandleUserNew(ctx context.Context, us UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user data.User
		// TODO: validation
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		userPublic, err := us.New(ctx, user)
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusInternalServerError)
			return
		}
		enc := json.NewEncoder(w)
		enc.Encode(userPublic)
		return
	}
}

func HandleUserDelete(ctx context.Context, us UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idRaw := r.PathValue("id")
		id, err := strconv.Atoi(idRaw)
		// TODO: validation
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusBadRequest)
			return
		}

		err = us.Delete(ctx, id)
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		return
	}
}
