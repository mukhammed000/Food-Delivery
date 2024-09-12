package main

import (
	"auth/api"
	"auth/api/handler"
	"auth/config"
	"auth/service"
	"auth/storage/postgres"
	"log"

	_ "auth/docs"
)

func main() {
	cfg := config.Load()

	// Initialize Postgres storage
	postgresStorage, err := postgres.NewPostgresStorage(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Printf("Database connected successfully! %+v\n", postgresStorage)

	// Initialize AuthService and Handler
	authService := service.NewAuthService(postgresStorage)
	apiHandler := handler.NewHandler(authService)

	// Initialize and run Gin router
	router := api.NewGin(apiHandler)
	if err := router.Run(cfg.HTTP_PORT); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
