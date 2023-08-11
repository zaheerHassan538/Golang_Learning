package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Wg is used to wait for the program to finsih

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

//main is the entry point for all Go Programs

func main() {
	// create an unbuffered channel
	court := make(chan int)

	// Add a count of two, one for each goroutine
	wg.Add(2)

	// Launch two players
	go player("Nadal", court)
	go player("Djokovic", court)

	// start the set
	court <- 1

	// Wait for the game to finsih
	wg.Wait()

}

// player simulates a person playing the tennis

func player(name string, court chan int) {
	// schedule the call to done to tell main we are done
	defer wg.Done()

	for {
		// wait for the ball to be hit back to us
		ball, ok := <-court
		if !ok {
			// if channel closed we won
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// pick a random number and see if we miss the ball
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			// close the channel to signal we lost

			close(court)
			return
		}

		// Display and then increment the hit count by one

		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// HIt the ball back to the opposing player
		court <- ball

	}
}
