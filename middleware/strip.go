package middleware

import (
	"net/http"
)

func StripSlashes(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func RedirectSlashes(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func StripPrefix(prefix string) func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}
