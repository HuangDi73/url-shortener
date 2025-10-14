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

func NewHandler(mux *http.ServeMux, deps HandlerDeps) {
	h := handler(deps)
	mux.HandleFunc("POST /auth/login", h.Login())
	mux.HandleFunc("POST /auth/register", h.Register())
}

func (h *handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.ContentLength == 0 {
			http.Error(w, "Empty body", http.StatusBadRequest)
			return
		}
		body, err := req.HandleBody[LoginRequest](w, r)
		if err != nil {
			return
		}
		fmt.Println(*body)
		data := LoginResponse{
			Token: "asdafaf",
		}
		res.Json(w, data, http.StatusOK)
	}
}

func (h *handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
