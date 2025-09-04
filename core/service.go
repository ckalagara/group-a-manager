package core

import (
	"context"
	"log"
	"time"

	pb "github.com/ckalagara/group-a-manager/proto"
)

type ServerImpl struct {
	pb.UnimplementedServiceServer
}

func (s *ServerImpl) Health(ctx context.Context, req *pb.HealthRequest) (*pb.HealthResponse, error) {
	log.Printf("Received health request at %v", time.Now())
	return &pb.HealthResponse{Status: "Up"}, nil
}
