package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in golang")

	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "tomato"

	fruitList[3] = "peach"

	fmt.Println("Fruit list is", fruitList)
	fmt.Println("Length of fruit list", len(fruitList))

	var vegList = [3]string{"potato", "mushroom", "beans"}
	fmt.Println("Veg list is", vegList)
	fmt.Println("Length of veg list", len(vegList))

}
