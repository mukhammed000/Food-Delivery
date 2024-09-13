package main

import (
	"log"

	"api/api"
	"api/api/handler"
	pb "api/genproto/delivery"

	config "api/config"

	_ "api/docs"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg := config.Load()

	// Connecting to serivce
	deliveryConn, err := grpc.Dial(cfg.DELIVERY_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while creating new client of delivery service: ", err.Error())
	}
	defer deliveryConn.Close()

	csc := pb.NewCartServiceClient(deliveryConn)
	osc := pb.NewOrdersServiceClient(deliveryConn)
	oscc := pb.NewOfficeClient(deliveryConn)
	psc := pb.NewProductServiceClient(deliveryConn)

	// Create a new handler with the clients
	h := handler.NewHandler(csc, oscc, osc, psc)
	r := api.NewGin(h)

	// Start the Gin server
	err = r.Run(cfg.HTTP_PORT)
	if err != nil {
		log.Fatalln("Error while running server: ", err.Error())
	}
}
