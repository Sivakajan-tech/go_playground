package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v3"
)

type Comment struct {
	Text string `form:"text" json:"text"`
}

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	// Define a route group for API version 1
	api := app.Group("/api/v1")

	// Define a route for creating comments
	api.Post("/comments", createComment)

	// Start the server and listen on port 3000
	app.Listen(":3000")

}

func ConnectProducer(brokersUrl []string) (*sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		log.Println("Error while creating producer: ", err)
		return nil, err
	}

	return producer, nil
}

func PushCommentToQueue(topic string, comment []byte) error {
	brokersUrl := []string{"localhost:9092"}
	producer, err := ConnectProducer(brokersUrl)
	if err != nil {
		log.Println("Error while connecting to the producer: ", err)
		return err
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(comment),
	}

	partition, offset, err := (*producer).SendMessage(msg)
	if err != nil {
		log.Println("Error while sending message")
		return err
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	return nil
}

func createComment(c *fiber.Ctx) error {
	comment := new(Comment)

	if err := c.BodyParser(comment); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message":   "An error occurred while parsing the request body",
			"succeeded": false,
		})
	}

	cmtInBytes, err := json.Marshal(comment)
	PushCommentToQueue("comments", cmtInBytes)

	err = c.JSON(&fiber.Map{
		"succeeded": true,
		"message":   "Comment created successfully",
		"comment":   comment,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"succeeded": false,
			"message":   "An error occurred while processing the request",
		})
	}
	return err
}
