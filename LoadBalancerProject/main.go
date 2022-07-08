package main

import (
	"fmt"
	"net/http"

	loadbalancer "github.com/sky-2002/LoadBalancerProject/LoadBalancer"
	"github.com/sky-2002/LoadBalancerProject/servers"
)

func main() {
	my_servers := []servers.Server{
		servers.NewSimpleServer("https://www.google.com"),
		servers.NewSimpleServer("https://www.bing.com"),
		servers.NewSimpleServer("https://www.duckduckgo.com"),
	}

	lb := loadbalancer.NewLoadBalancer("8080", my_servers)
	lb.AddSimpleServer("https://www.facebook.com")
	fmt.Println("New server added")

	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.ServeProxy(rw, req)
	}

	// register a proxy handler to handle all requests
	http.HandleFunc("/c", handleRedirect) // as this address does not exist, you can see that sites change at each reload

	fmt.Printf("serving requests at 'localhost:%s'\n", lb.Port)
	http.ListenAndServe(":"+lb.Port, nil)
}
