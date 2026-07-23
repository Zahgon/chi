package middleware

import (
	"net/http"
)

func PageRoute(path string, handler http.Handler) func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}
