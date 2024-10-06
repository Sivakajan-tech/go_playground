package main

import (
	"io"
	"log"

	pb "github.com/Sivakajan-tech/go_playground/greeting_grpc/proto"
)

func (s *helloServer) SayHelloClientStream(stream pb.GreetService_SayHelloClientStreamServer, req *pb.NameList) error {
	var message []string
	for {
		msg, err := stream.Recv()
		if err != io.EOF {
			return stream.SendAndClose(&pb.MessageList{Messages: message})
		}
		if err != nil {
			return err
		}
		log.Printf("Got message %v", msg.Name)

		message = append(message, "Hello", msg.Name)
	}
}
