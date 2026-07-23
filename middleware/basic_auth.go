package middleware

import (
	"net/http"
)

func BasicAuth(realm string, creds map[string]string) func(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func basicAuthFailed(w http.ResponseWriter, realm string) { _ = "STUB: not implemented"; return }
