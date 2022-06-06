package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to channels in Go")

	myCh := make(chan int, 2) // create channel
	wg := &sync.WaitGroup{}

	// myCh <- 5 // push 5 in the channel
	// // use <- channel_name to receive values from the channel
	// fmt.Println(<-myCh)

	// ---------------IMPORTANT-------------------------------------------------------------------------
	// The above code leads to deadlock because channel can only accept values
	// if someone is listening to it. And it cannot pass a value into the channel
	// because the code that is listening to it is on line 12 and we are passing the value in line 10
	// --------------------------------------------------------------------------------------------------

	wg.Add(2)
	// this <- arrow makes it(the function) receive only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myCh

		// fmt.Println(<-myCh)
		fmt.Println(isChannelOpen, val)

		wg.Done()
	}(myCh, wg)

	// this <- arrow makes it(the function) send only
	go func(ch chan<- int, wg *sync.WaitGroup) {

		// myCh <- 5
		// myCh <- 15
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
