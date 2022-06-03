package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println("The time now is", presentTime)
	fmt.Println("The time now is", presentTime.Format("01-02-2006"))        // use this syntax to get mm/dd/yy format
	fmt.Println("The time now is", presentTime.Format("01-02-2006 Monday")) // use this exactly to get the day as well
	fmt.Println("The time now is", presentTime.Format("02-01-2006"))        // this has dd/mm/yy format
	// also we can add time in between (as given in documentation)

	createdDate := time.Date(2020, time.August, 15, 23, 23, 0, 0, time.UTC)
	fmt.Println("Created date is", createdDate.Format("02-01-2006"))
}
