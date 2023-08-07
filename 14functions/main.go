package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is :", result)

	proRes := proAdder(2, 5, 8, 7)
	fmt.Println("Pro result is :", proRes)
}

func greeter() {
	fmt.Println("good for goolang")
}

// you are not allowed write function inside a function

func adder(valone int, valTwo int) int {
	return valone + valTwo

}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}
