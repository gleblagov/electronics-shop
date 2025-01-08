package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gleblagov/electronics-shop/data"
	"github.com/gleblagov/electronics-shop/utils"
)

type UserStorage interface {
	GetById(ctx context.Context, id int) (data.UserPublic, error)
	New(ctx context.Context, user data.User) (data.UserPublic, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, newBody data.User) (data.UserPublic, error)
	GetByEmail(ctx context.Context, email string) (data.User, error)
}

func HandleUserLogin(ctx context.Context, us UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		creds := struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}{}
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusBadRequest)
			return
		}
		user, err := us.GetByEmail(ctx, creds.Email)
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusBadRequest)
			return
		}
		if !utils.VerifyPass(creds.Password, user.Password) {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}
		token, err := utils.GenerateJwt(user.Email, user.Role)
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   token,
			Expires: time.Now().Add(24 * time.Hour),
		})
	}
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

func HandleUserUpdate(ctx context.Context, us UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idRaw := r.PathValue("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusBadRequest)
			return
		}

		var body data.User
		// TODO: validation
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		userPublic, err := us.Update(ctx, id, body)
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusInternalServerError)
			return
		}
		enc := json.NewEncoder(w)
		enc.Encode(userPublic)
		return
	}
}
