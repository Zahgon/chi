package chi

import "net/http"

func Chain(middlewares ...func(http.Handler) http.Handler) Middlewares {
	_ = "STUB: not implemented"
	return *new(Middlewares)
}

func (mws Middlewares) Handler(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (mws Middlewares) HandlerFunc(h http.HandlerFunc) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

type ChainHandler struct {
	Endpoint    http.Handler
	chain       http.Handler
	Middlewares Middlewares
}

func (c *ChainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func chain(middlewares []func(http.Handler) http.Handler, endpoint http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}
