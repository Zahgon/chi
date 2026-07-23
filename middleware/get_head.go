package middleware

import (
	"net/http"
)

func GetHead(next http.Handler) http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }
