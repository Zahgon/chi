package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5/_examples/versions/data"
)

type Article struct {
	*data.Article

	Data map[string]bool `json:"data" xml:"data"`
}

func (a *Article) Render(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func NewArticleResponse(article *data.Article) *Article { _ = "STUB: not implemented"; return nil }
