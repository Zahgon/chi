package middleware

import "net/http"

func New(h http.Handler) func(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

type contextKey struct {
	name string
}

func (k *contextKey) String() string { _ = "STUB: not implemented"; return "" }
