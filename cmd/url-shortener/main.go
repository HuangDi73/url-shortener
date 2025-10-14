package main

import (
	"log"
	"net/http"
	"url-shortener/config"
)

func main() {
	conf := config.Load()
	mux := http.NewServeMux()

	server := http.Server{
		Addr:    conf.Port,
		Handler: mux,
	}

	log.Printf("Server is running on http://localhost%s", server.Addr)
	server.ListenAndServe()
}
