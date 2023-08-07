package main

import "fmt"

type user struct {
	name  string
	email string
}

// notify implements a method with a value reciever
func (u user) notify() {
	fmt.Printf("Sending User Email to %s<%s>\n", u.name, u.email)

}

// changeEmail implements a method with a pointer reciever
func (u *user) changeEmail(email string) {
	u.email = email
}

// main is the entry point for the application

func main() {
	//values of type user can be used to call methods
	//declared with a value receiver

	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	//pointers of type user can also be used to call methods
	//declared a with a value reciever

	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()

	// value of type user can be used to call methods
	//declared with a point receiver

	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	lisa.changeEmail("lisa@comcast.com")
	lisa.notify()
}
