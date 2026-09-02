package middleware

import (
	"net/http"
)

func ContentCharset(charsets ...string) func(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func contentEncoding(ce string, charsets ...string) bool { _ = "STUB: not implemented"; return false }

func split(str, sep string) (string, string) { _ = "STUB: not implemented"; return "", "" }
