package link

import (
	"net/http"
	"url-shortener/config"
)

type handler struct {
	*config.Config
}

type HandlerDeps struct {
	*config.Config
}

func NewHandler(mux *http.ServeMux, deps HandlerDeps) {
	h := handler(deps)
	mux.HandleFunc("POST /link", h.Create())
	mux.HandleFunc("GET /{hash}", h.GoTo())
	mux.HandleFunc("PATCH /link/{id}", h.Update())
	mux.HandleFunc("DELETE /link/{id}", h.Delete())
}

func (h *handler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *handler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
