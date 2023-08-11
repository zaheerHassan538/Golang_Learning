package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	// counter is a variable incremented by all goroutines
	counter int64
	// wg  is used to wait for the program to finish
	wg sync.WaitGroup
)

// main is the entry point of all go programs
func main() {
	// add a count of 2 , one for each goroutine
	wg.Add(2)

	// create two goriutines
	go incCounter(1)
	go incCounter(2)

	// wait for the goroutines to finish
	wg.Wait()
	fmt.Println("Final Counter:", counter)

}

// incounter increments the package level  counter variable
func incCounter(id int) {
	// schedule the call to Done to tell main we are done
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// capture the value of counter
		//value := counter

		// safely add one to counter
		atomic.AddInt64(&counter, 1)
		//yield the thread and be palced back in queue
		runtime.Gosched()

		// Increment our local value of counter
		//value++

		// store the value back into counter
		//counter = value
	}

}
