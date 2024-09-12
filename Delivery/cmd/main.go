package main

import (
	"delivery/config"
	"delivery/genproto/delivery"
	"delivery/service"
	"delivery/storage/postgres"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	// Storage
	stg, err := postgres.NewPostgresStorage(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connected successfully!")

	// Connection
	lis, err := net.Listen("tcp", cfg.HTTP_PORT)
	if err != nil {
		log.Fatalf("Error while creating TCP listener: %v", err)
	}
	defer lis.Close()

	server := grpc.NewServer()
	service := service.NewDeliveryService(stg)

	// Registering services
	delivery.RegisterCartServiceServer(server, service)
	delivery.RegisterOfficeServer(server, service)
	delivery.RegisterOrdersServiceServer(server, service)
	delivery.RegisterProductServiceServer(server, service)

	// Run
	log.Println("Server listening at", cfg.HTTP_PORT)
	if server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
