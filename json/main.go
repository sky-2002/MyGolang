package main

import (
	"encoding/json"
	"fmt"
)

type Point struct {
	X int
	Y int
}

type Employee struct {
	Name    string
	Age     int
	Address string
}

func main() {
	fmt.Println("Exploring JSON in golang (the very basics)")

	p1 := Point{X: 2, Y: 5}
	jp, _ := json.Marshal(p1)
	fmt.Println(string(jp))
	var pt Point
	json.Unmarshal(jp, &pt)
	fmt.Println(pt)

	fmt.Println("")
	emp := Employee{Name: "George Smith", Age: 30, Address: "Newyork, USA"}
	empData, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(empData))

	// var name string = "Aakash"
	// var tf bool = true
	// var it int = 10

	// var uname string
	// var utf bool
	// var uit int

	// // json.Marshal returns byte,error
	// jname, _ := json.Marshal(name) // converting to bytes
	// fmt.Println(jname)             // bytes
	// fmt.Println(string(jname))     // converting bytes to string
	// json.Unmarshal(jname, &uname)
	// fmt.Println(uname)
	// fmt.Println("")

	// jtf, _ := json.Marshal(tf) // converting to bytes
	// fmt.Println(jtf)           // bytes
	// fmt.Println(string(jtf))   // converting bytes to string
	// json.Unmarshal(jtf, &utf)
	// fmt.Println(utf)
	// fmt.Println("")

	// jit, _ := json.Marshal(it) // converting to bytes
	// fmt.Println(jit)           // bytes
	// fmt.Println(string(jit))   // converting bytes to string
	// json.Unmarshal(jit, &uit)
	// fmt.Println(uit)
	// fmt.Println("")
}
