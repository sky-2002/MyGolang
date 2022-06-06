package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to racing in Go")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	var score = []int{0}

	wg.Add(4)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Routine 1")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Routine 2")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Routine 3")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	// use Read lock
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Read Lock")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()

	// fmt.Println(score)
}
