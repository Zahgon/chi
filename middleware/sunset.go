package middleware

import (
	"net/http"
	"time"
)

func Sunset(sunsetAt time.Time, links ...string) func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}
