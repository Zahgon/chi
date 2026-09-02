package middleware

import (
	"net/http"
)

var trueClientIP = http.CanonicalHeaderKey("True-Client-IP")
var xForwardedFor = http.CanonicalHeaderKey("X-Forwarded-For")
var xRealIP = http.CanonicalHeaderKey("X-Real-IP")

func RealIP(h http.Handler) http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

func realIP(r *http.Request) string { _ = "STUB: not implemented"; return "" }
