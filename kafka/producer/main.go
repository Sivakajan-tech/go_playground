package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v3"
)

// Comment struct represents a comment with a text field.
// The `form` tag is used for parsing form data, while the `json` tag is used for parsing JSON data.
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
	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Error starting server: ", err)
	}

}

// ConnectProducer creates a new synchronous Kafka producer.
// It takes a slice of broker URLs as input and returns a sarama.SyncProducer
// and an error. The function configures the producer to wait for all acknowledgments
// and allows for a maximum of 5 retries on failure. If the producer cannot be created,
// it logs the error and returns nil along with the error.
func ConnectProducer(brokersUrl []string) (sarama.SyncProducer, error) {
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

// PushCommentToQueue sends a comment to a specified Kafka topic.
// This function establishes a connection to the Kafka producer, constructs a message,
// and sends it to the specified topic. It logs any errors encountered during the process.
func PushCommentToQueue(topic string, comment []byte) error {
	/* Parameters:
	   - topic: A string representing the Kafka topic to which the comment will be sent.
	   - comment: A byte slice containing the comment data to be sent.
	Returns:
	   - An error if there is an issue connecting to the producer or sending the message; otherwise, it returns nil.
	*/
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

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Println("Error while sending message")
		return err
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	return nil
}

// createComment handles the creation of a new comment.
// It reads the JSON request body, unmarshals it into a Comment struct,
// pushes the comment to a Kafka queue, and responds with the created comment.
// In case of any errors during these processes, it returns appropriate error messages.
func createComment(c fiber.Ctx) error {
	comment := new(Comment)

	// Read the request body
	body := c.Body()
	// Unmarshal the JSON body into the Comment struct
	if err := json.Unmarshal(body, comment); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message":   "An error occurred while parsing the request body",
			"succeeded": false,
		})
	}

	// Marshal the Comment struct back to JSON bytes
	cmtInBytes, err := json.Marshal(comment)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"succeeded": false,
			"message":   "An error occurred while marshaling the comment",
		})
	}

	// Push the comment to the Kafka queue
	if err := PushCommentToQueue("comments", cmtInBytes); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"succeeded": false,
			"message":   "An error occurred while pushing the comment to the queue",
		})
	}

	// Respond with success message and the created comment
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
