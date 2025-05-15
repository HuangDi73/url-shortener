package auth

import (
	"fmt"
	"net/http"
	"url-shortener/config"
	"url-shortener/pkg/req"
	"url-shortener/pkg/res"
)

type handler struct {
	*config.Config
}

type HandlerDeps struct {
	*config.Config
}

func NewHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := handler(deps)
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (h handler) Login() http.HandlerFunc {
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

func (h handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
