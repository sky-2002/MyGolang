package main

import "fmt"

type Point struct {
	x int32
	y int32
}

type Circle struct {
	radius float64
	center *Point
}

// func

func main() {
	p1 := Point{x: 1, y: 2}
	fmt.Println(p1)

	p2 := Point{1, 3}
	fmt.Println(p2)

	c1 := Circle{radius: 5, center: &p1}
	fmt.Println(c1)
	fmt.Println(c1.center.x, c1.center.y)
	fmt.Println(c1.radius)
}
