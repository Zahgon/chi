package v3

import (
	"net/http"

	"github.com/go-chi/chi/v5/_examples/versions/data"
)

type Article struct {
	*data.Article `json:",inline" xml:",inline"`

	URL        string `json:"url" xml:"url"`
	ViewsCount int64  `json:"views_count" xml:"views_count"`
	APIVersion string `json:"api_version" xml:"api_version"`

	CustomDataForAuthUsers interface{} `json:"custom_data,omitempty" xml:"custom_data,omitempty"`
}

func (a *Article) Render(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func NewArticleResponse(article *data.Article) *Article { _ = "STUB: not implemented"; return nil }
