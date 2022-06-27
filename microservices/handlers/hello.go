package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// to create a handler, create a struct which implements the
// interface http.handler

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello world !")

	// reading from the user
	data, err := ioutil.ReadAll(r.Body)
	// log.Printf("Data %s\n", data)
	if err != nil {
		http.Error(rw, "Ooops", http.StatusBadRequest)
		return
	}

	// giving response to user
	fmt.Fprintf(rw, "Hello %s\n", data)
}
