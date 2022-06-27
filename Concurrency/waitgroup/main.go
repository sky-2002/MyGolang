package main

import (
	"fmt"
	"sync"
)

func main() {
	var beeper sync.WaitGroup

	ninjas := []string{"Tommy", "John", "Adam", "James"}
	beeper.Add(len(ninjas))

	for _, n := range ninjas {
		go attack(n, &beeper)
	}
	beeper.Wait()
	fmt.Println("Mission completed")
}

func attack(ninja string, beeper *sync.WaitGroup) {
	fmt.Println("Attacked", ninja)
	beeper.Done()
}
