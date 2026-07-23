package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/docgen"
	"github.com/go-chi/render"
)

var routes = flag.Bool("routes", false, "Generate router documentation")

func main() {
	flag.Parse()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})

	r.Route("/articles", func(r chi.Router) {
		r.With(paginate).Get("/", ListArticles)
		r.Post("/", CreateArticle)
		r.Get("/search", SearchArticles)

		r.Route("/{articleID}", func(r chi.Router) {
			r.Use(ArticleCtx)
			r.Get("/", GetArticle)
			r.Put("/", UpdateArticle)
			r.Delete("/", DeleteArticle)
		})

		r.With(ArticleCtx).Get("/{articleSlug:[a-z-]+}", GetArticle)
	})

	r.Mount("/admin", adminRouter())

	if *routes {

		fmt.Println(docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
			ProjectPath: "github.com/go-chi/chi/v5",
			Intro:       "Welcome to the chi/_examples/rest generated docs.",
		}))
		return
	}

	http.ListenAndServe(":3333", r)
}

func ListArticles(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func ArticleCtx(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func SearchArticles(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func CreateArticle(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func GetArticle(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func UpdateArticle(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func DeleteArticle(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func adminRouter() chi.Router { _ = "STUB: not implemented"; return *new(chi.Router) }

func AdminOnly(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func paginate(next http.Handler) http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

func init() {
	render.Respond = func(w http.ResponseWriter, r *http.Request, v interface{}) {
		if err, ok := v.(error); ok {

			if _, ok := r.Context().Value(render.StatusCtxKey).(int); !ok {
				w.WriteHeader(400)
			}

			fmt.Printf("Logging err: %s\n", err.Error())

			render.DefaultResponder(w, r, render.M{"status": "error"})
			return
		}

		render.DefaultResponder(w, r, v)
	}
}

type UserPayload struct {
	*User
	Role string `json:"role"`
}

func NewUserPayloadResponse(user *User) *UserPayload { _ = "STUB: not implemented"; return nil }

func (u *UserPayload) Bind(r *http.Request) error { _ = "STUB: not implemented"; return nil }

func (u *UserPayload) Render(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

type ArticleRequest struct {
	*Article

	User *UserPayload `json:"user,omitempty"`

	ProtectedID string `json:"id"`
}

func (a *ArticleRequest) Bind(r *http.Request) error { _ = "STUB: not implemented"; return nil }

type ArticleResponse struct {
	*Article

	User *UserPayload `json:"user,omitempty"`

	Elapsed int64 `json:"elapsed"`
}

func NewArticleResponse(article *Article) *ArticleResponse { _ = "STUB: not implemented"; return nil }

func (rd *ArticleResponse) Render(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func NewArticleListResponse(articles []*Article) []render.Renderer {
	_ = "STUB: not implemented"
	return nil
}

type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	StatusText string `json:"status"`
	AppCode    int64  `json:"code,omitempty"`
	ErrorText  string `json:"error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	_ = "STUB: not implemented"
	return *new(render.Renderer)
}

func ErrRender(err error) render.Renderer { _ = "STUB: not implemented"; return *new(render.Renderer) }

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Article struct {
	ID     string `json:"id"`
	UserID int64  `json:"user_id"`
	Title  string `json:"title"`
	Slug   string `json:"slug"`
}

var articles = []*Article{
	{ID: "1", UserID: 100, Title: "Hi", Slug: "hi"},
	{ID: "2", UserID: 200, Title: "sup", Slug: "sup"},
	{ID: "3", UserID: 300, Title: "alo", Slug: "alo"},
	{ID: "4", UserID: 400, Title: "bonjour", Slug: "bonjour"},
	{ID: "5", UserID: 500, Title: "whats up", Slug: "whats-up"},
}

var users = []*User{
	{ID: 100, Name: "Peter"},
	{ID: 200, Name: "Julia"},
}

func dbNewArticle(article *Article) (string, error) { _ = "STUB: not implemented"; return "", nil }

func dbGetArticle(id string) (*Article, error) { _ = "STUB: not implemented"; return nil, nil }

func dbGetArticleBySlug(slug string) (*Article, error) { _ = "STUB: not implemented"; return nil, nil }

func dbUpdateArticle(id string, article *Article) (*Article, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dbRemoveArticle(id string) (*Article, error) { _ = "STUB: not implemented"; return nil, nil }

func dbGetUser(id int64) (*User, error) { _ = "STUB: not implemented"; return nil, nil }
