package main

import (
	"log/slog"
	"net/http"
	"sync"
)

func main() {
	mux := http.NewServeMux()
	slog.Info("starting server...")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
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
