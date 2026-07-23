package middleware

import (
	"net/http"
)

var (
	URLFormatCtxKey = &contextKey{"URLFormat"}
)

func URLFormat(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}
