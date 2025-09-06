package core

import (
	"context"
	"log"
	"time"

	pbInventory "github.com/ckalagara/group-a-inventory/proto"
	pb "github.com/ckalagara/group-a-manager/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	inventoryUrl = "group-a-inventory:50052"
)

type ServerImpl struct {
	pb.UnimplementedServiceServer
}

func (s *ServerImpl) Health(ctx context.Context, req *pb.HealthRequest) (*pb.HealthResponse, error) {
	log.Printf("Received health request at %v", time.Now())
	// Ping inventory health endpoint
	dailOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	c, err := grpc.NewClient(inventoryUrl, dailOpts...)
	defer c.Close()
	if err != nil {
		log.Printf("Could not connect to inventory service: %v", err)
		return &pb.HealthResponse{Status: "Down"}, nil
	}

	invClient := pbInventory.NewServiceClient(c)
	pbRes, err := invClient.Health(ctx, &pbInventory.HealthRequest{})
	if err != nil {
		log.Printf("Inventory service health check failed: %v", err)
		return &pb.HealthResponse{Status: "Down"}, nil
	}
	log.Println("Inventory service is responding with status:", pbRes.Status)

	if pbRes.Status != "Service is healthy" {
		return &pb.HealthResponse{Status: "Down"}, nil
	}

	return &pb.HealthResponse{Status: "Up"}, nil
}
