# Use officical image for Golang
FROM golang:1.16.3-alpine3.13

# Set working directory
WORKDIR /app

# Copy all files from current directory to /app in container
COPY . .

# Download all dependencies
RUN go get -d -v ./...

# Build the application
RUN go bu-ild -o main .

# Expose port 8080 to the outside world
EXPOSE 8000

# Command to run the application
CMD [ "go run" "main.go"]