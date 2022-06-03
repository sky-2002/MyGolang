package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to maps in golang")
	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"

	fmt.Println("List of all languages", languages)
	fmt.Println("js:", languages["js"])

	// delete some values
	delete(languages, "rb")
	fmt.Println(languages)

	// for loop
	for key, value := range languages {
		fmt.Printf("For key %v value is %v\n", key, value)
	}
}
