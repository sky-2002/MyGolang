package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Monday", "Tueday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println(days)

	for _, day := range days {
		fmt.Println("Day is:", day)
	}

	for w := 0; w < 10; {
		fmt.Println(w)
		if w == 5 {
			goto here
		}
		w++
	}

here:
	fmt.Println("Reached at here (goto)")

}
