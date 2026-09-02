package chi

import (
	"net/http"
	"sync"
)

var _ Router = &Mux{}

type Mux struct {
	handler http.Handler

	tree *node

	methodNotAllowedHandler http.HandlerFunc

	parent *Mux

	pool *sync.Pool

	notFoundHandler http.HandlerFunc

	middlewares []func(http.Handler) http.Handler

	inline bool
}

func NewMux() *Mux { _ = "STUB: not implemented"; return nil }

func (mx *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func (mx *Mux) Use(middlewares ...func(http.Handler) http.Handler) {
	_ = "STUB: not implemented"
	return
}

func (mx *Mux) Handle(pattern string, handler http.Handler) { _ = "STUB: not implemented"; return }

func (mx *Mux) HandleFunc(pattern string, handlerFn http.HandlerFunc) {
	_ = "STUB: not implemented"
	return
}

func (mx *Mux) Method(method, pattern string, handler http.Handler) {
	_ = "STUB: not implemented"
	return
}

func (mx *Mux) MethodFunc(method, pattern string, handlerFn http.HandlerFunc) {
	_ = "STUB: not implemented"
	return
}

func (mx *Mux) Connect(pattern string, handlerFn http.HandlerFunc) {
	_ = "STUB: not implemented"
	return
}

func (mx *Mux) Delete(pattern string, handlerFn http.HandlerFunc) {
	_ = "STUB: not implemented"
	return
}

func (mx *Mux) Get(pattern string, handlerFn http.HandlerFunc) { _ = "STUB: not implemented"; return }

func (mx *Mux) Head(pattern string, handlerFn http.HandlerFunc) { _ = "STUB: not implemented"; return }

func (mx *Mux) Options(pattern string, handlerFn http.HandlerFunc) {
	_ = "STUB: not implemented"
	return
}

func (mx *Mux) Patch(pattern string, handlerFn http.HandlerFunc) { _ = "STUB: not implemented"; return }

func (mx *Mux) Post(pattern string, handlerFn http.HandlerFunc) { _ = "STUB: not implemented"; return }

func (mx *Mux) Put(pattern string, handlerFn http.HandlerFunc) { _ = "STUB: not implemented"; return }

func (mx *Mux) Query(pattern string, handlerFn http.HandlerFunc) { _ = "STUB: not implemented"; return }

func (mx *Mux) Trace(pattern string, handlerFn http.HandlerFunc) { _ = "STUB: not implemented"; return }

func (mx *Mux) NotFound(handlerFn http.HandlerFunc) { _ = "STUB: not implemented"; return }

func (mx *Mux) MethodNotAllowed(handlerFn http.HandlerFunc) { _ = "STUB: not implemented"; return }

func (mx *Mux) With(middlewares ...func(http.Handler) http.Handler) Router {
	_ = "STUB: not implemented"
	return *new(Router)
}

func (mx *Mux) Group(fn func(r Router)) Router { _ = "STUB: not implemented"; return *new(Router) }

func (mx *Mux) Route(pattern string, fn func(r Router)) Router {
	_ = "STUB: not implemented"
	return *new(Router)
}

func (mx *Mux) Mount(pattern string, handler http.Handler) { _ = "STUB: not implemented"; return }

func (mx *Mux) Routes() []Route { _ = "STUB: not implemented"; return nil }

func (mx *Mux) Middlewares() Middlewares { _ = "STUB: not implemented"; return *new(Middlewares) }

func (mx *Mux) Match(rctx *Context, method, path string) bool {
	_ = "STUB: not implemented"
	return false
}

func (mx *Mux) Find(rctx *Context, method, path string) string {
	_ = "STUB: not implemented"
	return ""
}

func (mx *Mux) NotFoundHandler() http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func (mx *Mux) MethodNotAllowedHandler(methodsAllowed ...methodTyp) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func (mx *Mux) handle(method methodTyp, pattern string, handler http.Handler) *node {
	_ = "STUB: not implemented"
	return nil
}

func (mx *Mux) routeHTTP(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func (mx *Mux) nextRoutePath(rctx *Context) string { _ = "STUB: not implemented"; return "" }

func (mx *Mux) updateSubRoutes(fn func(subMux *Mux)) { _ = "STUB: not implemented"; return }

func (mx *Mux) updateRouteHandler() { _ = "STUB: not implemented"; return }

func methodNotAllowedHandler(methodsAllowed ...methodTyp) func(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return nil
}
