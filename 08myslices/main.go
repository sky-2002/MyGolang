package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices tutorial")

	// for slices, we do not need to mention size
	var fruitList = []string{"apple", "tomato", "peach"}
	fmt.Printf("Type of fruit list is %T\n", fruitList)

	fruitList = append(fruitList, "mango", "banana")
	fmt.Println("Updated fruit list is", fruitList)

	// here we slice it from 1 to end, skipping 0th element
	// end at 3 but not including 3, similar to python
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777

	// append relallocates memory
	highScores = append(highScores, 555, 666, 321)

	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//  removing a value from a slice using index
	courses := []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
