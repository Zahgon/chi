package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/v3", func(r chi.Router) {
		r.Use(apiVersionCtx("v3"))
		r.Mount("/articles", articleRouter())
	})

	r.Route("/v2", func(r chi.Router) {
		r.Use(apiVersionCtx("v2"))
		r.Mount("/articles", articleRouter())
	})

	r.Route("/v1", func(r chi.Router) {
		r.Use(randomErrorMiddleware)
		r.Use(apiVersionCtx("v1"))
		r.Mount("/articles", articleRouter())
	})

	http.ListenAndServe(":3333", r)
}

func apiVersionCtx(version string) func(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func articleRouter() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

func listArticles(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func getArticle(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func randomErrorMiddleware(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}
