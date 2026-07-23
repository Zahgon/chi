package middleware

import (
	"net/http"
)

func WithValue(key, val interface{}) func(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}
