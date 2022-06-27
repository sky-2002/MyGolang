package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	ninja := "tommy"

	finished := make(chan bool) // a boolean channel, we will put a true in it once finished

	go attack(ninja, finished)
	// time.Sleep(time.Second) // total time spent was 3 sec
	// this shows that goroutines run in parallel/concurrently

	// ninjas := []string{"aakash", "messi", "ronaldo", "neymar", "suarez", "mbappe"}
	// start := time.Now()
	defer func() {
		fmt.Printf("Time spent: %s\n", time.Since(start))
	}()

	// fmt.Println("The channel contains: ", <-finished) // if this line is commente, it finishes in 0 seconds

	// attacking one by one
	// for _, ninja := range ninjas {
	// 	finished <- false // set channel as false, it will become true on every iteration, but this results in deadlock
	// 	// try using many channels
	// 	attack(ninja, finished)
	// }
	// this took approx 4 seconds

	// attacking
	// for _, ninja := range ninjas {
	// 	go attack(ninja) // spawning different processes
	// }
	// this took approx 2 seconds including the 2 second sleep at the end

	// time.Sleep(2 * time.Second)
}

func attack(target string, attacked chan bool) {
	time.Sleep(time.Second)
	fmt.Printf("Attacking target: %s\n", target)
	// time.Sleep(time.Second * 2) // time.Second is one second
	attacked <- true
}
