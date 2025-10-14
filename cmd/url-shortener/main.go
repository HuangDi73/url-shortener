package main

import (
	"log"
	"net/http"
	"url-shortener/config"
	"url-shortener/internal/auth"
	"url-shortener/internal/link"
)

func main() {
	conf := config.Load()
	mux := http.NewServeMux()

	// Handlers
	auth.NewHandler(mux, auth.HandlerDeps{Config: conf})
	link.NewHandler(mux, link.HandlerDeps{Config: conf})

	server := http.Server{
		Addr:    conf.Port,
		Handler: mux,
	}

	log.Printf("Server is running on http://localhost%s", server.Addr)
	server.ListenAndServe()
}
