package main

import (
	"context"
	"log/slog"
	"net/http"
	"sync"

	"github.com/gleblagov/electronics-shop/data"
	"github.com/gleblagov/electronics-shop/handlers"
)

func main() {
	ps, err := data.NewPostgresStorage()
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	slog.Info("starting server...")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		mux.HandleFunc("GET /user/{id}", handlers.HandleUserGetById(context.TODO(), ps.Users))
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
			panic(err)
		}
		wg.Done()
	}()
	slog.Info("server started")
	wg.Wait()
}
