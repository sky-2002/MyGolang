package main

import (
	"fmt"
	"time"
)

func main() {
	// goroutine is created by simply adding a keyword go before the statement
	go greeting("Aakash")
	greeting("Thatte")
}

func greeting(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}
}
