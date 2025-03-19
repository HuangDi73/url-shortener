package link

import (
	"net/http"
	"url-shortener/pkg/req"
	"url-shortener/pkg/res"
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
		body, err := req.HandleBody[CreateRequest](w, r)
		if err != nil {
			return
		}
		link := NewLink(body.Url)
		createdLink, err := h.Repo.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		res.Json(w, createdLink, http.StatusCreated)
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
