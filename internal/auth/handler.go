package auth

import (
	"net/http"
	"url-shortener/config"
	"url-shortener/pkg/jwt"
	"url-shortener/pkg/req"
	"url-shortener/pkg/res"
)

type handler struct {
	*config.Config
	AuthService *Service
}

type HandlerDeps struct {
	*config.Config
	AuthService *Service
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
		email, err := h.AuthService.Login(
			body.Email,
			body.Password,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		token, err := jwt.NewJWT(h.Auth.Secret).Create(email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := LoginResponse{
			Token: token,
		}
		res.Json(w, data, http.StatusOK)
	}
}

func (h *handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.ContentLength == 0 {
			http.Error(w, "Empty body", http.StatusBadRequest)
			return
		}
		body, err := req.HandleBody[RegisterRequest](w, r)
		if err != nil {
			return
		}
		email, err := h.AuthService.Register(
			body.Email,
			body.Name,
			body.Password,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		token, err := jwt.NewJWT(h.Auth.Secret).Create(email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := RegisterResponse{
			Token: token,
		}
		res.Json(w, data, http.StatusCreated)
	}
}
