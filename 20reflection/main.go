package main

import (
	"fmt"
	"reflect"
)

type details struct {
	fname   string
	lname   string
	age     int
	balance float64
}

type myType string

func showDetails(i, j interface{}) {
	t1 := reflect.TypeOf(i)
	k1 := t1.Kind()
	t2 := reflect.TypeOf(j)
	k2 := t2.Kind()
	fmt.Println("Type of first interface:", t1)
	fmt.Println("Kind of first interface:", k1)
	fmt.Println("Type of second interface", t2)
	fmt.Println("Kind of second interface", k2)

	fmt.Println("The values in the first arguments are :")
	if reflect.ValueOf(i).Kind() == reflect.Struct {
		value := reflect.ValueOf(i)
		numberofFields := value.NumField()
		for i := 0; i < numberofFields; i++ {
			fmt.Printf("%d.Type:%T || Value:%#v\n",
				(i + 1), value.Field(i), value.Field(i))
			fmt.Println("Kind is", value.Field(i).Kind())

		}

	}
	value := reflect.ValueOf(j)
	fmt.Printf("The Value passes in"+"second parameter is %#v", value)
}
func main() {

	iD := myType("12345678")
	person := details{
		fname:   "Go",
		lname:   "Geek",
		age:     32,
		balance: 33000.203,
	}
	showDetails(person, iD)

}

//func observe(i interface{}) {
//	fmt.Printf("The type passed is: %T\n", i)
// check the value in interface

//	fmt.Printf("The value passed is: %#v \n", i)
//	fmt.Println("----------------------------")

//}

//func main() {
//	var value float64 = 15
//	value2 := "GeeksForGeeks"
//	observe(value)
//	observe(value2)
//}
