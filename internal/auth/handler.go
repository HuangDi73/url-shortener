package auth

import (
	"encoding/json"
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
		data := LoginResponse{
			Token: "123",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}
}

func (h Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
