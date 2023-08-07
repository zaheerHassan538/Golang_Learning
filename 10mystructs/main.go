package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	//no inheritance in golang; no super or parent

	person := User{"Hassan", "hassanzah@gmial.com", true, 27}
	fmt.Println(person)
	fmt.Printf("hassan details are: %+v\n", person)
	fmt.Printf("Name is %v and email is %v.", person.Name, person.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
