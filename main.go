package main

import (
	"log"
	"net"

	"github.com/harlitad/notitication-service/service"

	"google.golang.org/grpc"
)

func main() {
	// Start the gRPC server
	grpcServer := grpc.NewServer()

	// call notification server
	service.NewNotificationServer(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	log.Println("Starting server on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
