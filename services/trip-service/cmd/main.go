package main

import (
	"context"
	"log"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
)

func main() {
	ctx := context.Background()

	inmemRepo := repository.NewInMemoryRepository()
	srv := service.NewService(inmemRepo)

	fare := &domain.RideFareModel{
		UserID:            "42",
		PackageSlug:       "luxury",
		TotalPriceInCents: 1000,
	}

	t, err := srv.CreateTrip(ctx, fare)
	if err != nil {
		log.Printf("Error creating trip: %v", err)
	}

	log.Printf("Trip created: %+v", t)

}
