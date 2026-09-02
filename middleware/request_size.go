package middleware

import (
	"net/http"
)

func RequestSize(bytes int64) func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}
