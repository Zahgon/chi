package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/users/{userID}", pathValueHandler)

	http.ListenAndServe(":3333", r)
}

func pathValueHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }
