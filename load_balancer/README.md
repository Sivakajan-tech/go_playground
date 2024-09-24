# Go Load Balancer

This project implements a simple Load Balancer in Go, which forwards HTTP requests to a pool of backend servers using the round-robin algorithm. The load balancer uses Go's `net/http` package and the `httputil.ReverseProxy` for proxying requests to backend servers.

## Features

- **Round-Robin Algorithm**: Distributes incoming requests across the backend servers in a round-robin manner.
- **Health Check (Stub)**: Each server is assumed to always be alive with a simple check (`IsAlive` method).
- **Reverse Proxy**: Proxies incoming HTTP requests to one of the available servers.
- **Custom Load Balancer**: Custom implementation of a load balancer that can be expanded to include health checks, server health monitoring, etc.

## Prerequisites

  - Go (v1.18+)
  - Internet access for the server URLs to be reachable


Here is a README.md file for your Go-based Load Balancer project:

markdown
Copy code
# Go Load Balancer

This project implements a simple Load Balancer in Go, which forwards HTTP requests to a pool of backend servers using the round-robin algorithm. The load balancer uses Go's `net/http` package and the `httputil.ReverseProxy` for proxying requests to backend servers.

## Features

- **Round-Robin Algorithm**: Distributes incoming requests across the backend servers in a round-robin manner.
- **Health Check (Stub)**: Each server is assumed to always be alive with a simple check (`IsAlive` method).
- **Reverse Proxy**: Proxies incoming HTTP requests to one of the available servers.
- **Custom Load Balancer**: Custom implementation of a load balancer that can be expanded to include health checks, server health monitoring, etc.

##  Installation

### 1. Clone the repository:
```bash
git clone https://github.com/Sivakajan-tech/go_playground.git
cd load_balancer
```

### 2. Run the project:
```bash
go run .
```
The load balancer will start listening on `localhost:8000`.

## Code Explanation

  - **Server Interface:** The `Server` interface defines three methods: `Address()`, `IsAlive()`, and `Serve()`.
  
  - **SimpleServer:** The `simpleServer` struct implements the `Server` interface. It holds the server's address and a reverse proxy to forward requests.

  - **LoadBalancer:** The `LoadBalancer` struct maintains a list of backend servers and forwards requests to them using a round-robin algorithm.
 
  - **Round-Robin Algorithm:** The `getNextAvailableServer` method ensures that requests are distributed evenly across servers.
  