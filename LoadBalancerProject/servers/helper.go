package servers

import (
	"fmt"
	"net/http/httputil"
	"net/url"

	errors "github.com/sky-2002/LoadBalancerProject/errors"
)

func NewSimpleServer(addr string) *simpleServer {
	serverUrl, err := url.Parse(addr)
	errors.HandleErr(err)
	fmt.Println("serverUrl:", serverUrl)
	return &simpleServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}
