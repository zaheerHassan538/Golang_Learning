package main

import (
	"fmt"
	"sync"
)

func runner1(wg *sync.WaitGroup) {
	defer wg.Done() // this decreases counter by 1
	fmt.Print("\nI am first runner")

}

func runner2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("\nI am second runner")

}

func execute() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	// we increasing the counter by 2 beacuse we have 2 GoRoutines
	go runner1(wg)
	go runner2(wg)

	// this blocks the execution untill its counter become zero
	wg.Wait()
}
func main() {
	// lunching both the runners
	execute()
}
