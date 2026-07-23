package middleware

import (
	"net/http"
)

func AllowContentEncoding(contentEncoding ...string) func(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}
