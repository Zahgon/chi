package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})

	r.Group(func(r chi.Router) {

		r.Use(middleware.Timeout(2500 * time.Millisecond))

		r.Get("/slow", func(w http.ResponseWriter, r *http.Request) {
			rand.Seed(time.Now().Unix())

			processTime := time.Duration(rand.Intn(4)+1) * time.Second

			select {
			case <-r.Context().Done():
				return

			case <-time.After(processTime):

			}

			w.Write([]byte(fmt.Sprintf("Processed in %v seconds\n", processTime)))
		})
	})

	r.Group(func(r chi.Router) {

		r.Use(middleware.Timeout(30 * time.Second))

		r.Use(middleware.Throttle(1))

		r.Get("/throttled", func(w http.ResponseWriter, r *http.Request) {
			select {
			case <-r.Context().Done():
				switch r.Context().Err() {
				case context.DeadlineExceeded:
					w.WriteHeader(504)
					w.Write([]byte("Processing too slow\n"))
				default:
					w.Write([]byte("Canceled\n"))
				}
				return

			case <-time.After(5 * time.Second):

			}

			w.Write([]byte("Processed\n"))
		})
	})

	http.ListenAndServe(":3333", r)
}
