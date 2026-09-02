package middleware

import (
	"net/http"
	"time"
)

func Timeout(timeout time.Duration) func(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}
