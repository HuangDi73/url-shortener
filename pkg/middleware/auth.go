package middleware

import (
	"context"
	"net/http"
	"strings"
	"url-shortener/config"
	"url-shortener/pkg/jwt"
)

type key string

const (
	CtxEmailKey key = "CtxEmailKey"
)

func writeUnauthed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthed(next http.Handler, conf *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authStr := r.Header.Get("Authorization")
		if !strings.HasPrefix(authStr, "Bearer ") {
			writeUnauthed(w)
			return
		}
		token := strings.TrimPrefix(authStr, "Bearer ")
		isValid, data := jwt.NewJWT(conf.Auth.Secret).Parse(token)
		if !isValid {
			writeUnauthed(w)
			return
		}
		ctx := context.WithValue(r.Context(), CtxEmailKey, data.Email)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
