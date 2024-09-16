# Kafka Comment Producer/Consumer API

This project demonstrates a simple API for creating comments and pushing them to a Kafka queue using a Kafka producer, as well as consuming these comments using a Kafka consumer.

## Getting Started

To get started, clone the repository:

```bash
git clone https://github.com/Sivakajan-tech/go_playground.git
cd kafka
```

## Prerequisites

Make sure you have the following installed:

- Go (1.16 or above)
- Kafka running locally or in your cloud environment.
  - Default configuration assumes Kafka is running on localhost:9092

Make sure to run the Kafka cluster and Zookeeper using either docker-compose or manually.

## Start application

Run both `consumer/main.go` and `producer/main.go` in separate terminals to achieve message production and consumption in the Kafka system.

```bash
go run consumer/main.go
```

```bash
go run producer/main.go
```

## Usage

### API Endpoints

- <b>POST /api/v1/comments</b><br>
  Create a comment and push it to the Kafka topic comments.

### Example Request

```bash
curl -X POST http://localhost:3000/api/v1/comments \
     -H "Content-Type: application/json" \
     -d '{"text": "This is a test comment"}'
```

### Response

```bash
{
  "succeeded": true,
  "message": "Comment created successfully",
  "comment": 
  {
    "text": "This is a test comment"
  }
}
```

### Kafka Consumer

To consume the messages pushed to the Kafka topic `comments`, run the Kafka consumer.

The consumer will print the consumed messages in the following format:

```bash
Received message count: 1 | Topic(comments) | Messages(This is a test comment)
```

**Note:** You can use Postman as well to send/receive requests to the API.
