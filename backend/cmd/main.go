package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"sync"

	"github.com/gleblagov/electronics-shop/data"
	"github.com/gleblagov/electronics-shop/handlers"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	ps, err := data.NewPostgresStorage()
	if err != nil {
		slog.Error("Failed to initialize PostgresStorage")
		panic(err)
	}

	mux := http.NewServeMux()
	slog.Info("Starting server...")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		mux.HandleFunc("GET /user/{id}", handlers.HandleUserGetById(context.TODO(), ps.Users))
		mux.HandleFunc("POST /user/login", handlers.HandleUserLogin(context.TODO(), ps.Users))
		mux.HandleFunc("POST /user", handlers.HandleUserNew(context.TODO(), ps.Users))
		mux.HandleFunc("DELETE /user/{id}", handlers.HandleUserDelete(context.TODO(), ps.Users))
		mux.HandleFunc("PATCH /user/{id}", handlers.HandleUserUpdate(context.TODO(), ps.Users))

		mux.HandleFunc("GET /product/{id}", handlers.HandleProductGetById(context.TODO(), ps.Products))
		mux.HandleFunc("POST /product", handlers.HandleProductNew(context.TODO(), ps.Products))
		mux.HandleFunc("DELETE /product/{id}", handlers.HandleProductDelete(context.TODO(), ps.Products))

		mux.HandleFunc("GET /cart/{id}", handlers.HandleCartGetById(context.TODO(), ps.Carts))
		mux.HandleFunc("POST /cart", handlers.HandleCartNew(context.TODO(), ps.Carts))
		mux.HandleFunc("PATCH /cart/{id}", handlers.HandleCartChangeStatus(context.TODO(), ps.Carts))
		mux.HandleFunc("POST /cart/{id}/product", handlers.HandleCartAddProduct(context.TODO(), ps.Carts))

		mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
			return
		})
		if err := http.ListenAndServe("0.0.0.0:3737", mux); err != nil {
			slog.Error("Failed to listen on 0.0.0.0:3737")
			panic(err)
		}
		wg.Done()
	}()
	slog.Info("Server started")
	wg.Wait()
}
