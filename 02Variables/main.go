package main

import "fmt"

const LoginToken string = "asfafsdfsd545645asd" // public

func main() {
	var username string = "Aakash"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username) // %T is for type, %v for values

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloatVal float32 = 255.65457798
	fmt.Println(smallFloatVal)
	fmt.Printf("Variable is of type: %T \n", smallFloatVal)

	var largeFloatVal float64 = 255.65457798
	fmt.Println(largeFloatVal)
	fmt.Printf("Variable is of type: %T \n", largeFloatVal)

	// default values and aliases
	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Variable is of type: %T \n", anotherVar)

	var anotherStr string
	fmt.Println(anotherStr)
	fmt.Printf("Variable is of type: %T \n", anotherStr)

	// implicit type
	var website = "aakash@go.com"
	fmt.Println(website)

	// no var style
	myNumber := 25
	fmt.Println(myNumber)

	fmt.Println("Login token is", LoginToken)
}
