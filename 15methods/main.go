package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	//no inheritance in golang; no super or parent

	person := User{"Hassan", "hassanzah@gmial.com", true, 27}
	fmt.Println(person)
	fmt.Printf("hassan details are: %+v\n", person)
	fmt.Printf("Name is %v and email is %v.", person.Name, person.Email)
	person.GetStatus()
	person.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user Active :", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}
