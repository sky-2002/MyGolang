package loadbalancer

import (
	"fmt"
	"net/http"

	"github.com/sky-2002/LoadBalancerProject/servers"
)

type LoadBalancer struct {
	Port            string
	roundRobinCount int
	AllServers      []servers.Server
}

func NewLoadBalancer(port string, servers []servers.Server) *LoadBalancer {
	return &LoadBalancer{
		Port:            port,
		roundRobinCount: 0,
		AllServers:      servers,
	}
}

// getNextServerAddr returns the address of the next available server to send a
// request to, using a simple round-robin algorithm
func (lb *LoadBalancer) getNextAvailableServer() servers.Server {
	server := lb.AllServers[lb.roundRobinCount%len(lb.AllServers)]
	for !server.IsAlive() {
		lb.roundRobinCount++
		server = lb.AllServers[lb.roundRobinCount%len(lb.AllServers)]
	}
	lb.roundRobinCount++

	return server
}

func (lb *LoadBalancer) ServeProxy(rw http.ResponseWriter, req *http.Request) {
	targetServer := lb.getNextAvailableServer()

	// could optionally log stuff about the request here!
	fmt.Printf("forwarding request to address %q\n", targetServer.Address())

	// could delete pre-existing X-Forwarded-For header to prevent IP spoofing
	targetServer.Serve(rw, req)
}

func (lb *LoadBalancer) AddSimpleServer(addr string) {
	lb.AllServers = append(lb.AllServers, servers.NewSimpleServer(addr))
}
