package main

import "fmt"

func main() {

	// Using var
	// var name = "Batman"
	var age int32 = 37
	const isCool = true

	// shorthand
	name := "Batman"
	size := 1.7 // default type: float64
	email := "batman@gmail.com"

	name1, email1 := "Robin", "robin@gmail.com"

	// print to screen
	fmt.Println(name, email, age, isCool)

	fmt.Println(name1, email1)

	// get variable's type
	fmt.Printf("%T", size)
}
