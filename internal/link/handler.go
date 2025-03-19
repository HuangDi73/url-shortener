package link

import (
	"net/http"
)

type Handler struct {
	Repo *Repository
}

type HandlerDeps struct {
	Repo *Repository
}

func NewHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := Handler{
		Repo: deps.Repo,
	}
	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("GET /{hash}", handler.GoTo())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
}

func (h Handler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h Handler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h Handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h Handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
