package main

import (
	"sync"
)

func main() {
	// var beeper sync.WaitGroup

	// ninjas := []string{"Tommy", "John", "Adam", "James"}
	// beeper.Add(len(ninjas))

	// for _, n := range ninjas {
	// 	go attack(n, &beeper)
	// }
	// beeper.Wait()
	// fmt.Println("Mission completed")

	var beeper sync.WaitGroup
	beeper.Add(1)
	beeper.Done() // without this, it will result in a deadlock
	// beeper.Done() // adding this twice gives a negative waitgroup counter, then add 2 and it will be ok
	beeper.Wait()

}

// func attack(ninja string, beeper *sync.WaitGroup) {
// 	fmt.Println("Attacked", ninja)
// 	beeper.Done()
// }
