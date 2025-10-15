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
	"url-shortener/pkg/event"
	"url-shortener/pkg/middleware"
)

func main() {
	conf := config.Load()
	db := db.NewDb(conf)
	mux := http.NewServeMux()
	eventBus := event.NewEventBus()

	// Repositories
	linkRepo := link.NewRepository(db)
	userRepo := user.NewRepository(db)
	statRepo := stat.NewRepository(db)

	// Services
	authService := auth.NewService(userRepo)
	statService := stat.NewService(&stat.ServiceDeps{
		EventBus: eventBus,
		StatRepo: statRepo,
	})

	// Handlers
	auth.NewHandler(mux, auth.HandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewHandler(mux, link.HandlerDeps{
		Config:   conf,
		LinkRepo: linkRepo,
		EventBus: eventBus,
	})
	stat.NewHandler(mux, stat.HandlerDeps{
		StatRepo: statRepo,
		Config:   conf,
	})

	stack := middleware.Chain(
		middleware.Logging,
		middleware.Cors,
	)

	server := http.Server{
		Addr:    conf.Port,
		Handler: stack(mux),
	}

	go statService.AddClick()

	log.Printf("Server is running on http://localhost%s", server.Addr)
	server.ListenAndServe()
}
