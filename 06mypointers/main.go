package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers tutorial")

	var one int = 2
	var ptr *int = &one
	fmt.Println("value of pointer is", ptr)
	fmt.Println("value present at reference", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value at reference", *ptr)
	fmt.Println("Value of one", one)

	// var dptr *int
	// fmt.Println("value of pointer is", dptr) // default value of pointer is nil

}
