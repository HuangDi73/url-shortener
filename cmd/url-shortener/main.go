package main

import (
	"log"
	"net/http"
	"url-shortener/config"
	"url-shortener/internal/auth"
	"url-shortener/internal/link"
	"url-shortener/pkg/db"
)

func main() {
	conf := config.Load()
	db := db.NewDb(conf)
	mux := http.NewServeMux()

	// Repositories
	linkRepo := link.NewRepository(db)

	// Handlers
	auth.NewHandler(mux, auth.HandlerDeps{Config: conf})
	link.NewHandler(mux, link.HandlerDeps{
		Config: conf,
		Repo:   linkRepo,
	})

	server := http.Server{
		Addr:    conf.Port,
		Handler: mux,
	}

	log.Printf("Server is running on http://localhost%s", server.Addr)
	server.ListenAndServe()
}
