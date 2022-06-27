package main

import (
	"fmt"
	"time"
)

func main() {
	ninjas := []string{"aakash", "messi", "ronaldo", "neymar", "suarez", "mbappe"}
	start := time.Now()
	defer func() {
		fmt.Printf("Time spent: %s\n", time.Since(start))
	}()

	// attacking one by one
	// for _, ninja := range ninjas {
	// 	attack(ninja)
	// }
	// this took approx 4 seconds

	// attacking
	for _, ninja := range ninjas {
		go attack(ninja) // spawning different processes
	}
	// this took approx 2 seconds including the 2 second sleep at the end

	time.Sleep(2 * time.Second)
}

func attack(target string) {
	fmt.Printf("Attacking target: %s\n", target)
	time.Sleep(time.Second) // time.Second is one second
}
