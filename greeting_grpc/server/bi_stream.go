package main

import (
	"io"
	"log"

	pb "github.com/Sivakajan-tech/go_playground/greeting_grpc/proto"
)

func (s *helloServer) SayHelloBiDiStream(stream pb.GreetService_SayHelloBiDiStreamServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name : %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
