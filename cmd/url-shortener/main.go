package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	router := http.NewServeMux()

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: router,
	}

	log.Printf("Server is running in http://localhost%s", server.Addr)
	server.ListenAndServe()
}
