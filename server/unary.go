package main

import (
	"context"

	pb "github.com/nimamiri9248/grpc_sample/proto"
)

func (s *HelloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
