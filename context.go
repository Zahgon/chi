package chi

import (
	"context"
	"net/http"
)

func URLParam(r *http.Request, key string) string { _ = "STUB: not implemented"; return "" }

func URLParamFromCtx(ctx context.Context, key string) string { _ = "STUB: not implemented"; return "" }

func RouteContext(ctx context.Context) *Context { _ = "STUB: not implemented"; return nil }

func NewRouteContext() *Context { _ = "STUB: not implemented"; return nil }

var (
	RouteCtxKey = &contextKey{"RouteContext"}
)

type Context struct {
	Routes Routes

	parentCtx context.Context

	RoutePath   string
	RouteMethod string

	URLParams RouteParams

	routeParams RouteParams

	routePattern string

	RoutePatterns []string

	methodsAllowed   []methodTyp
	methodNotAllowed bool
}

func (x *Context) Reset() { _ = "STUB: not implemented"; return }

func (x *Context) Clone() *Context { _ = "STUB: not implemented"; return nil }

func (x *Context) URLParam(key string) string { _ = "STUB: not implemented"; return "" }

func (x *Context) RoutePattern() string { _ = "STUB: not implemented"; return "" }

func replaceWildcards(p string) string { _ = "STUB: not implemented"; return "" }

type RouteParams struct {
	Keys, Values []string
}

func (s *RouteParams) Add(key, value string) { _ = "STUB: not implemented"; return }

type contextKey struct {
	name string
}

func (k *contextKey) String() string { _ = "STUB: not implemented"; return "" }
