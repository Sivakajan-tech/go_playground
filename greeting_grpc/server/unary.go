package main

import (
	"context"

	pb "github.com/Sivakajan-tech/go_playground/greeting_grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello World"}, nil
}
