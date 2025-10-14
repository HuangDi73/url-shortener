package main

import (
	"log"
	"net/http"
	"url-shortener/config"
	"url-shortener/internal/auth"
	"url-shortener/internal/link"
	"url-shortener/internal/user"
	"url-shortener/pkg/db"
	"url-shortener/pkg/middleware"
)

func main() {
	conf := config.Load()
	db := db.NewDb(conf)
	mux := http.NewServeMux()

	// Repositories
	linkRepo := link.NewRepository(db)
	userRepo := user.NewRepository(db)

	// Services
	authService := auth.NewService(userRepo)

	// Handlers
	auth.NewHandler(mux, auth.HandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewHandler(mux, link.HandlerDeps{
		Config: conf,
		Repo:   linkRepo,
	})

	stack := middleware.Chain(
		middleware.Logging,
		middleware.IsAuthed,
		middleware.Cors,
	)

	server := http.Server{
		Addr:    conf.Port,
		Handler: stack(mux),
	}

	log.Printf("Server is running on http://localhost%s", server.Addr)
	server.ListenAndServe()
}
