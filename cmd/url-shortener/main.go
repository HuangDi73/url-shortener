package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	server := http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	log.Printf("Server is running on http://localhost%s", server.Addr)
	server.ListenAndServe()
}
