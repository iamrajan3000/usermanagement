package main

import (
	"log"
	"net"

	pb "Assignment_TotalityCorp/grpc"
	"Assignment_TotalityCorp/handlers"
	"Assignment_TotalityCorp/models"
	"Assignment_TotalityCorp/services"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	users := []*models.User{
		{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
		{Id: 2, Fname: "Alice", City: "NY", Phone: 9876543210, Height: 5.5, Married: false},
	}

	userService := services.NewUserService(users)
	userServer := handlers.NewUserServiceServer(userService)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userServer)

	log.Println("Starting gRPC server on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
