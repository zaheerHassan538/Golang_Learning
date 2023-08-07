package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch Case in golang")

	rand.Seed(time.now.UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("you can move 2 spot")
	case 3:
		fmt.Println("you can move 3 spot")
	case 4:
		fmt.Println("you can move 4 spot")
		fallthrough
	case 5:
		fmt.Println("you can move 5 spot")
	case 6:
		fmt.Println("you can move to 6 spot and roll dice again")
	default:
		fmt.Println("what was that!")
	}

}
