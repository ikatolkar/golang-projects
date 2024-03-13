package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func handleErr(err error) {
	if err != nil{
		fmt.Printf("error: %v\n", err)
	}
}

type Server interface {
	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, r *http.Request)
}

type simpleServer struct {
	addr string
	proxy *httputil.ReverseProxy
}

type LoadBalancer struct {
	port				string
	roundRobinCount		int
	servers				[]Server
}

/*
1. Create new Simple Server
2. Each server has an address (url)
3. And a reverse proxy to that address
*/
func newSimpleServer(addr string) *simpleServer {
	serverUrl, err := url.Parse(addr)
	handleErr(err)

	return &simpleServer{
		addr: addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

/*
1. Create a new load balancer, initialize
*/
func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port: 				port,
		roundRobinCount: 	0,
		servers: 			servers,
	}
}

/*
1. returns server's address
*/
func (s *simpleServer) Address() string {	return s.addr }

/*
1. Dummy function always returns true, ideally should check traffic level, availability, readiness etc
*/
func (s *simpleServer) IsAlive() bool { return true }

/*
1. Serve the request via the reverse proxy
*/
func (s *simpleServer) Serve(rw http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(rw, req)
}

/*
1. Pick a server
2. Check if the server is alive NOTE : IsAlive is a dummy function for now that always returns true
4. If found an alive server , increment round robin counter
5. Return the available server
*/
func (lb *LoadBalancer) getNextAvailableServer() Server {
	server := lb.servers[lb.roundRobinCount % len(lb.servers)]
	for !server.IsAlive()  {
		lb.roundRobinCount++
		server = lb.servers[lb.roundRobinCount % len(lb.servers)]
	}
	lb.roundRobinCount++
	return server
}

/*
1. Pick an available server in round robin
2. Start serving on that server
*/
func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, req *http.Request) {
	targetServer := lb.getNextAvailableServer()
	fmt.Printf("forwarding request to address %q\n", targetServer.Address())
	targetServer.Serve(rw, req)
}

/*
1. Start the reverse proxy server
2. Create a new load balancer
3. Handle routes at "/"
4. For every request at /, handle the redirect
*/
func main() {
	servers := []Server{
		newSimpleServer("https://www.facebook.com"),
		newSimpleServer("http://www.bing.com"),
		newSimpleServer("http://www.duckduckgo.com"),
	}
	lb := NewLoadBalancer("8000", servers)
	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.serveProxy(rw, req)
	}
	http.HandleFunc("/", handleRedirect)

	fmt.Printf("serving requests at 'localhost:%s'\n", lb.port)
	http.ListenAndServe(":"+lb.port, nil)
}
