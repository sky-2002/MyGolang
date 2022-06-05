package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Welcome to interfaces in golang")

	c1 := circle{radius: 5}
	r1 := rect{width: 2, length: 3}
	fmt.Printf("Area of circle is %v and of rectangle is %v\n", c1.area(), r1.area())

	shapes := []shape{c1, r1} // for this slice, we can only use area() function
	// because area() is the only common thing implemented by circle and rectangle
	fmt.Println(shapes)

	for _, s := range shapes {
		fmt.Println(s.area())
	}
}

type circle struct {
	radius float32
}
type rect struct {
	width  float32
	length float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}
func (r rect) area() float32 {
	return r.length * r.width
}

type shape interface {
	area() float32
}
