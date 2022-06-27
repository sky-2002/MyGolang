package main

import (
	"fmt"
)

func main() {
	n1, n2 := make(chan string), make(chan string)

	go Elect(n1, "Ninja 1")
	go Elect(n2, "Ninja 2")

	// fmt.Println(<-n1)
	// fmt.Println(<-n2)

	select {
	case message := <-n1:
		fmt.Println(message)
	case message := <-n2:
		fmt.Println(message)
	default:
		fmt.Println("Neither")
	}

	roughlyFair()

}

func Elect(n chan string, message string) {
	n <- message
}

func roughlyFair() {
	n1 := make(chan interface{})
	close(n1)
	n2 := make(chan interface{})
	close(n2)

	// go Elect(n1, "Ninja 1")
	// go Elect(n2, "Ninja 2")

	// fmt.Println(<-n1)
	// fmt.Println(<-n2)

	var n1c, n2c int
	for i := 0; i < 1000; i++ {
		select {
		case <-n1:
			// fmt.Println(message)
			n1c++
		case <-n2:
			// fmt.Println(message)
			n2c++
		}
	}
	fmt.Println("Ninja 1 count:", n1c)
	fmt.Println("Ninja 2 count:", n2c)
}
