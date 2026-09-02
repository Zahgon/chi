package data

import (
	"errors"
	"net/http"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrForbidden    = errors.New("Forbidden")
	ErrNotFound     = errors.New("Resource not found")
)

func PresentError(r *http.Request, err error) (*http.Request, interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}
