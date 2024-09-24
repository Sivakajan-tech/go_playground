package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Server interface {
	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, req *http.Request)
}

type simpleServer struct {
	address string
	proxy   *httputil.ReverseProxy
}

// Creates a new instance of simpleServer with the given address
func newSimpleServer(address string) *simpleServer {
	serverURL, err := url.Parse(address)
	handleError(err)

	return &simpleServer{
		address: address,
		proxy:   httputil.NewSingleHostReverseProxy(serverURL),
	}
}

// LoadBalancer struct holds the configuration for the load balancer
// port: the port on which the load balancer listens
// roundRobinCount: a counter to keep track of the next server to forward the request to
// servers: a list of servers that the load balancer can forward requests to
type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
}

func newLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobinCount: 0,
		servers:         servers,
	}
}

// Implementing methods of the Server interface
// Address returns the address with which to access the server
func (s *simpleServer) Address() string {
	return s.address
}

// IsAlive returns true if the server is alive and able to serve requests
func (s *simpleServer) IsAlive() bool {
	return true
}

// Serve uses this server to process the request
func (s *simpleServer) Serve(rw http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(rw, req)
}

/*
The function getNextAvailableServer is a receiver function (or method) in Go.
It is associated with the LoadBalancer type. Receiver functions allow you
to define methods on types, which can then be called on instances of those types.
*/

// getNextAvailableServer returns the next available server using round-robin algorithm
// This is a receiver function for the LoadBalancer type
func (lb *LoadBalancer) getNextAvailableServer() Server {
	server := lb.servers[lb.roundRobinCount%len(lb.servers)]
	for !server.IsAlive() {
		lb.roundRobinCount++
		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}
	lb.roundRobinCount++

	return server
}

// serverProxy forwards the request to the next available server using the round-robin algorithm
// This is a receiver function for the LoadBalancer type
func (lb *LoadBalancer) serverProxy(rw http.ResponseWriter, req *http.Request) {
	targetServer := lb.getNextAvailableServer()
	fmt.Printf("Redirecting request to server: %s\n", targetServer.Address())
	targetServer.Serve(rw, req)
}

func main() {
	servers := []Server{
		newSimpleServer("https://www.facebook.com"),
		newSimpleServer("https://www.bing.com"),
		newSimpleServer("https://www.duckduckgo.com"),
	}

	lb := newLoadBalancer("8000", servers)

	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.serverProxy(rw, req)
	}
	http.HandleFunc("/", handleRedirect)

	fmt.Printf("Load balancer listening on 'localhost%s'\n", lb.port)

	err := http.ListenAndServe(":"+lb.port, nil)
	handleError(err)
}
