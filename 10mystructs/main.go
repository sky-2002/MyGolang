package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in golang")

	// there is NO inheritance in golang

	aakash := User{"Aakash", "aakash@go.com", true, 20, 1234}
	fmt.Println(aakash)
	fmt.Printf("Aakash details are: %+v\n", aakash)
	fmt.Printf("Name is: %+v and Email is:%+v\n", aakash.Name, aakash.Email)
	fmt.Printf("Passcode is %+v\n", aakash.passcode)
}

// as we are using first letters capital, we can use these outside as well
type User struct {
	Name     string
	Email    string
	Status   bool
	Age      int
	passcode int // we can still use this, but we cannot export
}
