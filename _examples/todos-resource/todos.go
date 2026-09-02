package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type todosResource struct{}

func (rs todosResource) Routes() chi.Router { _ = "STUB: not implemented"; return *new(chi.Router) }

func (rs todosResource) List(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (rs todosResource) Create(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (rs todosResource) Get(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (rs todosResource) Update(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (rs todosResource) Delete(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (rs todosResource) Sync(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
