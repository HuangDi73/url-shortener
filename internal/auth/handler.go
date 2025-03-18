package auth

import (
	"net/http"
	"url-shortener/config"
)

type Handler struct {
	config.Config
}

type HandlerDeps struct {
	config.Config
}

func NewHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := Handler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (h Handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
