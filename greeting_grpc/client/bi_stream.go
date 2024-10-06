package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Sivakajan-tech/go_playground/greeting_grpc/proto"
)

func callSayHelloBiDiStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Bidirectional Streaming started")
	stream, err := client.SayHelloBiDiStream(context.Background())
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming %v", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		time.Sleep(2 * time.Second)
	}

	if err := stream.CloseSend(); err != nil {
		log.Fatalf("Error while closing the stream: %v", err)
	}
	<-waitc
	log.Printf("Bidirectional Streaming finished")
}
