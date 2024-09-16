package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

// main is the entry point of the Kafka consumer application.
// It connects to a Kafka broker, consumes messages from the specified topic,
// and handles incoming messages and errors. The application listens for
// termination signals (SIGINT, SIGTERM) to gracefully shut down the consumer.
// It counts and prints the number of messages processed before exiting.
func main() {
	topic := "comments"
	worker, err := connectConsumer([]string{"localhost:9092"})
	if err != nil {
		log.Println("Error while connecting to the consumer: ", err)
		return
	}

	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Println("Error while consuming partition: ", err)
		return
	}

	fmt.Println("Consumer started")
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	msgCount := 0
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			// Handle errors from the consumer
			case err := <-consumer.Errors():
				fmt.Println(err)
			// Handle incoming messages
			case msg := <-consumer.Messages():
				msgCount++
				fmt.Printf("Received message count: %d: | Topic(%s) | Messages(%s)\n", msgCount, string(msg.Topic), string(msg.Value))
			// Handle termination signals
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed", msgCount, "messages")
	if err := consumer.Close(); err != nil {
		log.Println("Error while closing the consumer: ", err)
	}
}

// connectConsumer establishes a connection to a Kafka consumer using the provided broker URLs.
// It returns a sarama.Consumer instance and an error if the connection fails.
func connectConsumer(brokersUrl []string) (sarama.Consumer, error) {
	/* 	Parameters:
	- brokersUrl: A slice of strings representing the URLs of the Kafka brokers to connect to.

		Returns:
	- A sarama.Consumer instance if the connection is successful.
	- An error if there was an issue creating the consumer.
	*/
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	conn, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		log.Println("Error while creating consumer: ", err)
		return nil, err
	}

	return conn, nil
}
