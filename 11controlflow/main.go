package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to control flow in golang")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "It is 10"
	}

	fmt.Println("The result is", result)
}
