package main

import (
	"fmt"
	"net/http"

	loadbalancer "github.com/sky-2002/LoadBalancerProject/LoadBalancer"
	"github.com/sky-2002/LoadBalancerProject/servers"
)

func main() {
	servers := []servers.Server{
		servers.NewSimpleServer("https://www.google.com"),
		servers.NewSimpleServer("https://www.bing.com"),
		servers.NewSimpleServer("https://www.duckduckgo.com"),
	}

	lb := loadbalancer.NewLoadBalancer("8080", servers)

	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.ServeProxy(rw, req)
	}

	// register a proxy handler to handle all requests
	http.HandleFunc("/", handleRedirect)

	fmt.Printf("serving requests at 'localhost:%s'\n", lb.Port)
	http.ListenAndServe(":"+lb.Port, nil)
}
