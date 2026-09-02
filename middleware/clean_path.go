package middleware

import (
	"net/http"
)

func CleanPath(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}
