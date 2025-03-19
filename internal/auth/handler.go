package auth

import (
	"fmt"
	"net/http"
	"url-shortener/config"
	"url-shortener/pkg/req"
	"url-shortener/pkg/res"
)

type Handler struct {
	*config.Config
}

type HandlerDeps struct {
	*config.Config
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
		payload, err := req.HandleBody[LoginRequest](w, r)
		if err != nil {
			return
		}
		fmt.Println(payload)
		data := LoginResponse{
			Token: "123",
		}
		res.Json(w, data, http.StatusOK)
	}
}

func (h Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
