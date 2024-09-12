# Use officical image for Golang
FROM golang:1.23.1

# Set working directory
WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and
# only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy all files from current directory to /app in container
COPY . .

# Build the application
RUN go build -v -o /usr/local/bin/app ./...

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the application
CMD ["go", "run", "main.go"]