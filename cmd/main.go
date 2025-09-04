package main

import (
	"log"
	"net"
	"time"

	"github.com/ckalagara/group-a-manager/core"
	pb "github.com/ckalagara/group-a-manager/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.Printf("Starting gRPC server on port 50051: %v", time.Now())
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen with error: %v", err)
	}
	log.Printf("Creating gRPC server: %v", time.Now())
	server := grpc.NewServer()
	reflection.Register(server)
	pb.RegisterServiceServer(server, &core.ServerImpl{})
	log.Printf("Serving gRPC server: %v", time.Now())

	err = server.Serve(listener)

	if err != nil {
		log.Fatalf("Failed to serve with error: %v", err)
	}

}
