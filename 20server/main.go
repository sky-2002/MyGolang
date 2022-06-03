package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web server in go")
	// PerformGetRequest()
	// PerformPostJSONRequest()
	PerformFormRequest()
}

func PerformGetRequest() {
	const myurl string = "http://localhost:8000/get" // we can also give this as a parameter to the function

	resp, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("Status code:", resp.StatusCode)
	fmt.Println("Content length:", resp.ContentLength)

	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(content))
}

func PerformPostJSONRequest() {
	const myurl string = "http://localhost:8000/post" // we can also give this as a parameter to the function

	// fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename":"Let's us go with golang",
			"price":"0",
			"platform":"learncodeonline.in"
		}
	`)

	resp, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(content))
}

func PerformFormRequest() {
	const myurl string = "http://localhost:8000/postform" // we can also give this as a parameter to the function

	// form data
	data := url.Values{}
	data.Add("firstname", "aakash")
	data.Add("lastname", "thatte")
	data.Add("email", "at@go.com")

	resp, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(content))
}
