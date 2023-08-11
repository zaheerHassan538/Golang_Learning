package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4  // no of goroutine to use
	taskLoad         = 10 // Amount of work to process
)

// wg is used to wait for the program to finish

var wg sync.WaitGroup

// init is called to intialiaze the package to the GO runtime prior to any other code being executed

func init() {
	// Seed the random number generator
	rand.Seed(time.Now().Unix())
}

// mainis the entry point for all go programs

func main() {
	// create a buffered channel to manage the task load

	tasks := make(chan string, taskLoad)

	// Lunch goroutines to handle the work
	wg.Add(numberGoroutines)

	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// Add a bunch of work to get done

	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	// close the channel so the goroutine will quit when all work is done
	wg.Wait()

}

// worker is lanuched as a goroutine to process work from the buffered channel

func worker(tasks chan string, worker int) {
	// report that we just returned
	defer wg.Done()

	for {
		// wait for work to be assigned
		task, ok := <-tasks
		if !ok {
			// This means the channel is empty and closed
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}

		// Display we are starting the work
		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		// Randomly wait to simulate work
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// Display we finsihed the work
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}
