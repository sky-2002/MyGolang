package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(3, 4)
	fmt.Println(result)

	fmt.Println(proAdder(1, 2, 3, 4))

}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(values ...int) int {
	total := 0
	for _, v := range values {
		total = total + v
	}
	return total
}

func greeter() {
	fmt.Println("Hello from greeter, Golang")
}
