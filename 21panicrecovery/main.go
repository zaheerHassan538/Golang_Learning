package main

import "fmt"

//func main() {
//	fmt.Println("Panic! in the Go App")
//	defer func() {
//		fmt.Println("I am a deffer function")
//	}()
//	panic("oh no")
//	fmt.Println("end of main")
//}

func sillySusan() {
	fmt.Println("silly susan called")
	panickingPeter()
	fmt.Println("silly susan finshed")
}

func panickingPeter() {
	fmt.Println("panicking peter called")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("hackers have been thwarted")
		}
	}()
	panic("oh no")
	fmt.Println("panicking peter finished")
}

func main() {
	fmt.Println("Cascading panics")
	sillySusan()
	fmt.Println("end of main func")
}
