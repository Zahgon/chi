package middleware

import (
	"net/http"
)

func RouteHeaders() HeaderRouter { _ = "STUB: not implemented"; return *new(HeaderRouter) }

type HeaderRouter map[string][]HeaderRoute

func (hr HeaderRouter) Route(header, match string, middlewareHandler func(next http.Handler) http.Handler) HeaderRouter {
	_ = "STUB: not implemented"
	return *new(HeaderRouter)
}

func (hr HeaderRouter) RouteAny(header string, match []string, middlewareHandler func(next http.Handler) http.Handler) HeaderRouter {
	_ = "STUB: not implemented"
	return *new(HeaderRouter)
}

func (hr HeaderRouter) RouteDefault(handler func(next http.Handler) http.Handler) HeaderRouter {
	_ = "STUB: not implemented"
	return *new(HeaderRouter)
}

func (hr HeaderRouter) Handler(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

type HeaderRoute struct {
	Middleware func(next http.Handler) http.Handler
	MatchOne   Pattern
	MatchAny   []Pattern
}

func (r HeaderRoute) IsMatch(value string) bool { _ = "STUB: not implemented"; return false }

type Pattern struct {
	prefix   string
	suffix   string
	wildcard bool
}

func NewPattern(value string) Pattern { _ = "STUB: not implemented"; return *new(Pattern) }

func (p Pattern) Match(v string) bool { _ = "STUB: not implemented"; return false }
