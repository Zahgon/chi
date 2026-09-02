package middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SupressNotFound(router *chi.Mux) func(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}
