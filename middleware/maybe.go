package middleware

import "net/http"

func Maybe(mw func(http.Handler) http.Handler, maybeFn func(r *http.Request) bool) func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}
