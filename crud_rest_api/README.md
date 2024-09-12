# CRUD Rest API

This repository contains a Go application that provides a basic RESTful API for managing users. The application connects to a MySQL database and is designed to help familiarize you with Golang and Docker technologies.

## Features

- Basic CRUD operations for managing users
- MySQL database integration
- Docker and Docker Compose setup for running the app in containers

## Requirements

- [Go](https://golang.org/doc/install) (for local development)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Getting Started

### Running Locally

If you want to run the application locally without Docker, ensure you have MySQL installed and running, and update the connection string in the `main.go` file.

1. Install Go modules:

   ```bash
   go mod download
   ```

2. Update the database connection string in `main.go` if necessary:

   ```go
   db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/<db_name>")
   ```

3. Run the Go application:
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`.

### Running with Docker Compose

You can run the entire application, including the MySQL database, using Docker Compose.

1. Build and start the application using Docker Compose:

   ```bash
   docker-compose up --build
   ```

2. Once everything is up, the Go application will be accessible at `http://localhost:8080`.

### API Endpoints

| Method | Endpoint      | Description             |
| ------ | ------------- | ----------------------- |
| GET    | `/users`      | Fetch all users         |
| GET    | `/users/{id}` | Fetch a user by ID      |
| POST   | `/users`      | Create a new user       |
| PUT    | `/users/{id}` | Update an existing user |
| DELETE | `/users/{id}` | Delete a user           |

### Environment Variables

The following environment variables are used to configure the database connection:

- `DB_HOST`: Hostname of the MySQL database (default: `db`)
- `DB_PORT`: MySQL port (default: `3306`)
- `DB_USER`: MySQL username (default: `root`)
- `DB_PASSWORD`: MySQL password (default: `password`)
- `DB_NAME`: MySQL database name (default: `go_playground`)

These are preconfigured in the `docker-compose.yml` file.

### Stopping the Application

To stop the application and remove the containers, run:

```bash
docker-compose down
```
