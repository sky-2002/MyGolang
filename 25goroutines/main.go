package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var WG sync.WaitGroup

func main() {
	// goroutine is created by simply adding a keyword go before the statement
	// go greeting("Aakash")
	// greeting("Thatte")

	websites := []string{
		"https://lco.dev",
		// "https://go.dec",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	t1 := time.Now()
	// without wait groups
	for _, v := range websites {
		getStatusCodeSimple(v)
	}
	t2 := time.Now()
	fmt.Println("Time:", t2.Sub(t1)) // this prints out 4.112 secongs
	fmt.Println("------------------")

	t1 = time.Now()
	// with wait groups
	for _, v := range websites {
		go getStatusCode(v)
		WG.Add(1)
	}
	WG.Wait() // this will not let the main method finish till all wait groups give it the signal to go ahead
	t2 = time.Now()
	fmt.Println("Time:", t2.Sub(t1)) // And this prints out 675.33 miliseconds, much faster
}

// func greeting(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(1 * time.Second)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {

	defer WG.Done() // wait group sends done message at the end of this function

	result, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Oops ! Error with endpoint")
	}
	fmt.Printf("%d status code for %s\n", result.StatusCode, endpoint)
}

func getStatusCodeSimple(endpoint string) {

	// defer WG.Done() // wait group sends done message at the end of this function

	result, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Oops ! Error with endpoint")
	}
	fmt.Printf("%d status code for %s\n", result.StatusCode, endpoint)
}
