package middleware

import (
	"net/http"
)

func SetHeader(key, value string) func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func AllowContentType(contentTypes ...string) func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}
