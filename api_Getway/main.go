package main

import (
	"getway/api"
	"log/slog"

	// _ "api-getway/docs"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up gRPC connections
	libraryConn, err := grpc.NewClient("lib-service-cont:50020", grpc.WithTransportCredentials(insecure.NewCredentials())) // Update the address
	if err != nil {
		slog.Error("Failed to connect to libraryConn service %v")
	}
	defer libraryConn.Close()

	router := api.NewGin(libraryConn)
	// fmt.Println("API Gateway running on http://localhost:50020")
	if err := router.Run(":8000"); err != nil {
		slog.Error("Failed to connect to gin engine: %v", "ERRORA", err)
	}

}
