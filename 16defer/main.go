package main

import "fmt"

func main() {
	fmt.Println("Welcome to defer")

	// this gets placed just before the last curly brace
	defer fmt.Println("tutorial")

	// if there many defer, they come in LIFO order

	fmt.Println("This is last line")

	myDefer()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
