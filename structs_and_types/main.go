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

type Student struct {
	name   string
	age    int
	grades []int
}

// -----------------------------------------
// getters
func (s Student) getGrades() []int {
	return s.grades
}

func (s Student) getName() string {
	return s.name
}

func (s Student) getAge() int { return s.age }

// -----------------------------------------

// setters
func (s *Student) setName(name string) {
	s.name = name
}
func (s *Student) setAge(age int) {
	s.age = age
}
func (s *Student) setGrades(grades []int) {
	s.grades = grades
}

//--------------------------------------------

func main() {
	// p1 := Point{x: 1, y: 2}
	// fmt.Println(p1)

	// p2 := Point{1, 3}
	// fmt.Println(p2)

	// c1 := Circle{radius: 5, center: &p1}
	// fmt.Println(c1)
	// fmt.Println(c1.center.x, c1.center.y)
	// fmt.Println(c1.radius)

	tim := Student{name: "Tim", grades: []int{70, 75, 85, 80}, age: 19}
	fmt.Println(tim)
	fmt.Println(tim.getName())
	fmt.Println(tim.getAge())
	fmt.Println(tim.getGrades())

	tim.setAge(20)
	tim.setName("Tom")
	tim.setGrades([]int{20, 20, 35, 36})

	fmt.Println(tim)
	fmt.Println(tim.getName())
	fmt.Println(tim.getAge())
	fmt.Println(tim.getGrades())

}
