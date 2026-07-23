package middleware

import (
	"net/http"
)

func Heartbeat(endpoint string) func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}
