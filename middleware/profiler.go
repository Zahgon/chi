//go:build !tinygo

package middleware

import (
	"net/http"
)

func Profiler() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }
