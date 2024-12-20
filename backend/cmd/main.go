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
