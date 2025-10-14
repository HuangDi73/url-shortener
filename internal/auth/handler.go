package auth

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
	mux.HandleFunc("POST /auth/login", h.Login())
	mux.HandleFunc("POST /auth/register", h.Register())
}

func (h *handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
