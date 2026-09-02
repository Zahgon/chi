package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "data"))
	FileServer(r, "/files", filesDir)

	http.ListenAndServe(":3333", r)
}

func FileServer(r chi.Router, path string, root http.FileSystem) { _ = "STUB: not implemented"; return }
