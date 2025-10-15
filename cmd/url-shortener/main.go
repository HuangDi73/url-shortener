package main

import (
	"log"
	"net/http"
	"url-shortener/config"
	"url-shortener/internal/auth"
	"url-shortener/internal/link"
	"url-shortener/internal/stat"
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
	statRepo := stat.NewRepository(db)

	// Services
	authService := auth.NewService(userRepo)

	// Handlers
	auth.NewHandler(mux, auth.HandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewHandler(mux, link.HandlerDeps{
		Config:   conf,
		LinkRepo: linkRepo,
		StatRepo: statRepo,
	})

	stack := middleware.Chain(
		middleware.Logging,
		middleware.Cors,
	)

	server := http.Server{
		Addr:    conf.Port,
		Handler: stack(mux),
	}

	log.Printf("Server is running on http://localhost%s", server.Addr)
	server.ListenAndServe()
}
