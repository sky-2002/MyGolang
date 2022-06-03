package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "https://lco.dev"

func main() {
	fmt.Println("Welcome to web requests in go")

	response, err := http.Get(URL)

	checkNilError(err)

	// the response type is a pointer
	// we get a reference of the response
	fmt.Printf("Response is of type %T\n", response)

	defer response.Body.Close() // it is our responsibility to close the connection

	// fmt.Println(response)
	databytes, err := ioutil.ReadAll(response.Body)
	checkNilError(err)
	fmt.Println("Databytes:", databytes)                   // gives the content in bytes
	fmt.Println("Databytes in string:", string(databytes)) // gives the html content
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
