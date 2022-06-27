package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rp http.ResponseWriter, r *http.Request) {
		log.Println("Hello world !")

		// reading from the user
		data, err := ioutil.ReadAll(r.Body)
		// log.Printf("Data %s\n", data)
		if err != nil {
			http.Error(rp, "Ooops", http.StatusBadRequest)
			return
		}

		// giving response to user
		fmt.Fprintf(rp, "Hello %s\n", data)
	})

	http.HandleFunc("/gb", func(http.ResponseWriter, *http.Request) {
		log.Println("GoodBye world !")
	})

	http.ListenAndServe(":9090", nil) // the DefaultServeMux is used, as we have passed nil
}
