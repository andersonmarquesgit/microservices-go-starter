package main

import (
	"log"
	"net/http"
	http2 "ride-sharing/services/trip-service/internal/infrastructure/http"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
)

func main() {
	inmemRepo := repository.NewInMemoryRepository()
	srv := service.NewService(inmemRepo)
	mux := http.NewServeMux()

	httpHandler := http2.HandlerService{Service: srv}

	mux.HandleFunc("POST /preview", httpHandler.HandleTripPreview)

	server := &http.Server{
		Addr:    ":8083",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("HTTP Server error: %v", err)
	}
}
