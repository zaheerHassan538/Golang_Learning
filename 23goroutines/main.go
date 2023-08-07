package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// allocate 1 logical processor for the scheduler to use

	runtime.GOMAXPROCS(1)

	// wg is used to wait for the program to finish
	// Add a count of two, one for each goroutine

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	//Declare an anonymous function and create a goroutine
	go func() {
		//schedule the call to done to tell main we are done
		defer wg.Done()

		//Display the alphabet 3 times
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// Declare an anonymous function and create a goroutine
	go func() {
		// schedule the call to DOne to tell main we are done
		defer wg.Done()

		//Display the alphabet 3 times
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	// wait for the goroutines to finsih
	fmt.Println("Waiting to Finsh")
	wg.Wait()

	fmt.Println("\n Terminating Program")
}
