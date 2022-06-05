package main

import "fmt"

func change1(str *string) {
	*str = "Changed it !"
}

func change2(str string) {
	str = "Changed it simply"
}

func main() {

	// here the value is 7 and the reference is the address where this value is stored
	x := 7
	// we use &x to get address/reference of x. &x is a pointer pointing to address of x.
	fmt.Println(&x)
	// we use * to get value at address
	fmt.Println(*&x) // this is same as printing x

	y := &x        // y is address of x
	*y = *y + 5    // increment the value at address y by 5
	fmt.Println(x) // This gives 12

	hello := "hello"
	fmt.Println(hello)
	change1(&hello)
	fmt.Println(hello)

	pointer := &hello
	fmt.Println("The pointer:", pointer)
	fmt.Println("Address of pointer itself:", &pointer)
}
