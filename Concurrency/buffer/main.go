package main

import "fmt"

func main() {
	// channel of name string
	// channel := make(chan string) // channel with capacity 0
	// these channels need two routines to exchange information

	channel := make(chan string, 3)
	// when used sequentially, a channel is a FIFO queue
	// setup goroutine
	func() {
		channel <- "First Message"
	}()

	func() {
		channel <- "Second Message"
	}()

	func() {
		channel <- "Third Message"
	}()

	fmt.Println("Channel:[", <-channel, "<-", <-channel, "<-", <-channel, "]")
	// fmt.Println("Channel:", <-channel)
}
