package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan string)

	go throwStar(channel)
	for {
		message, open := <-channel
		if open != true {
			break
		}
		fmt.Println(message)
	}
}

func throwStar(channel chan string) {
	rounds := 3
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < rounds; i++ {
		score := rand.Intn(10)

		channel <- fmt.Sprint("Score is:", score)
	}
	// Sprint returns the string
	close(channel) // this tells that this channel is no longer in use
}
