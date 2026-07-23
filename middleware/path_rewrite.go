package middleware

import (
	"net/http"
)

func PathRewrite(old, new string) func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}
