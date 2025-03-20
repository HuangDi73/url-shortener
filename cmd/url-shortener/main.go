package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"url-shortener/config"
	"url-shortener/internal/auth"
	"url-shortener/internal/link"
	"url-shortener/pkg/db"
	"url-shortener/pkg/middleware"
)

func main() {
	conf := config.Load()
	db := db.New(conf)
	router := http.NewServeMux()

	// Repositories
	linkRepo := link.NewRepository(db)

	// Handlers
	auth.NewHandler(router, auth.HandlerDeps{
		Config: conf,
	})
	link.NewHandler(router, link.HandlerDeps{
		Repo: linkRepo,
	})

	// Middlewares
	stack := middleware.Chain(
		middleware.Logging,
		middleware.CORS,
	)

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: stack(router),
	}

	log.Printf("Server is running in http://localhost%s", server.Addr)
	server.ListenAndServe()
}
