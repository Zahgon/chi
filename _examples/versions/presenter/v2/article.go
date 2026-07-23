package v2

import (
	"net/http"

	"github.com/go-chi/chi/v5/_examples/versions/data"
)

type Article struct {
	*data.Article

	SelfURL string `json:"self_url" xml:"self_url"`

	URL interface{} `json:"url,omitempty" xml:"url,omitempty"`
}

func (a *Article) Render(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func NewArticleResponse(article *data.Article) *Article { _ = "STUB: not implemented"; return nil }
