package main

import (
	greetings "command-line-arguments/home/emumba/mygolang-ico/16module/main.go"
	"fmt"
)

func main() {
	message := greetings.Hello("Hassan Zaheer")
	fmt.Println(message)
}
