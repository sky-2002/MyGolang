package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please give rating to pizza:")

	// comma ok syntax / error ok syntax
	// err is like catching an error
	input, err := reader.ReadString('\n') // read till there is a new line character in user's input
	fmt.Print("User entered:", input)
	fmt.Printf("Type of rating is %T", input)
	fmt.Println("Error is", err)
}
