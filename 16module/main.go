package greetings

import "fmt"

func Hello(name string) string {
	// := operator is shortcut for declaring and intializing a variable in one line
	// var message string
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
