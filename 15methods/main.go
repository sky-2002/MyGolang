package main

import "fmt"

func main() {
	// functions in structs are called methods
	fmt.Println("Welcome to methods in Go")

	u := User{"Aakash", "aakash@go.com", true, 20, 1234}
	u.GetStatus()
	u.NewMail("aakash@golang.com")
	fmt.Println("New email is", u.Email)
}

type User struct {
	Name     string
	Email    string
	Status   bool
	Age      int
	passcode int // we can still use this, but we cannot export
}

// this is a method
func (u User) GetStatus() {
	fmt.Println("The status of user is", u.Status)
}

// this user that we are passing, a copy of it is passed. Use pointer to pass actual object
// Now when we pass a reference to the object, the actual email gets updated.
func (u *User) NewMail(newemail string) {
	u.Email = newemail
	fmt.Println("Email updated !")
}
