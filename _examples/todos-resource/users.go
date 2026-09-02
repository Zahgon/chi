package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type usersResource struct{}

func (rs usersResource) Routes() chi.Router { _ = "STUB: not implemented"; return *new(chi.Router) }

func (rs usersResource) List(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (rs usersResource) Create(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (rs usersResource) Get(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (rs usersResource) Update(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (rs usersResource) Delete(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
