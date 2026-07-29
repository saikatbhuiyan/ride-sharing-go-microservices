package main

import (
	"context"
	"log"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"time"
)

func main() {
	ctx := context.Background()

	inmemRepo := repository.NewInMemoryRepository()

	svc := service.NewTripService(inmemRepo)

	fare := &domain.RideFareModel{
		UserID: primitive.NewObjectID() // example user ID
	}

	t, err := svc.CreateTrip(ctx, fare)
	if err != nil {
		log.Println(err)
	}

	log.Println(t)

	// keep the program running for now
	for {
		time.Sleep(time.Second)
	}
}
