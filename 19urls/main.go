package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghhtjkj"

func main() {
	fmt.Println("Welcome to urls in go")
	fmt.Println(myurl)

	// error
	result, _ := url.Parse(myurl)
	// fmt.Println(result.Scheme)   // gives https
	// fmt.Println(result.Host)     // gives lco.dev:3000
	// fmt.Println(result.Path)     // gives /learn
	// fmt.Println(result.Port())   // gives 3000
	// fmt.Println(result.RawQuery) // gives the params passed, coursename=reactjs?paymentid=ghhtjkj

	qparams := result.Query()
	fmt.Printf("The type of qparams is %T\n", qparams) // gives url.values
	// fmt.Println(qparams)
	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])

	for _, value := range qparams {
		fmt.Println("Param is:", value)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=james",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
