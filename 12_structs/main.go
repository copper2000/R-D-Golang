package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string

	// f, l, c, g string => you can declare like this with value has the same type
	age int
}

// Greeting method (value receiver)
func (person Person) greet() string {
	return "Hello, my name is " + person.firstName + " " + person.lastName + " and I am " +
		strconv.Itoa(person.age)
}

// HasBirthDay method (pointer receiver)
func (person *Person) hasBirthday() {
	person.age += 1
}

// GetMarried (poiner receiver)
func (person *Person) getMarried(spouseLastName string) {
	if person.gender == "male" {
		return
	} else {
		person.lastName = spouseLastName
	}
}

// Nếu không dùng con trỏ thì giá trị age chỉ được tăng trong hàm hasBirthday() => nên khi gọi hàm greet() giá trị của age vẫn là 32
// Nếu dùng con trỏ thì giá trị age sẽ bị thay đổi hoàn toàn vì con trỏ đã chọc đến địa chỉ của age => khi gọi hàm greet() giá trị của age sẽ là 33

func main() {
	// Init person using struct
	person1 := Person{
		firstName: "Jamie",
		lastName:  "Bell",
		city:      "Billingham",
		gender:    "Male",
		age:       34,
	}
	fmt.Println(person1)
	fmt.Println(person1.firstName, person1.age)

	// Alternative
	person2 := Person{"Kate", "Mara", "New York", "Female", 37}
	fmt.Println(person2)
	fmt.Println(person2.firstName, person2.age)

	fmt.Println("===================")
	// (value receiver)
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())

	fmt.Println("===================")
	// (pointer receiver)
	person1.hasBirthday()
	fmt.Println(person1.greet())

	fmt.Println("===================")
	// (pointer receiver)
	person2.hasBirthday()
	fmt.Println(person2.greet())

	// person1 married person2 => person2 is female so her last name gonna change by person1's lastname
	fmt.Println("Full Name before married:", person2.firstName+" "+person2.lastName)
	person2.getMarried(person1.lastName)
	fmt.Println("Last Name after married:", person2.firstName+" "+person2.lastName)

}
