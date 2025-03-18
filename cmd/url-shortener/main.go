package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"url-shortener/config"
	"url-shortener/internal/auth"
)

func main() {
	conf := config.Load()
	router := http.NewServeMux()
	auth.NewHandler(router, auth.HandlerDeps{Config: conf})

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: router,
	}

	log.Printf("Server is running in http://localhost%s", server.Addr)
	server.ListenAndServe()
}
