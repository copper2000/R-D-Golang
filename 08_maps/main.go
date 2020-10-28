package main

import "fmt"

func main() {
	// Define map
	emails := make(map[string]string)

	// Assign key value
	emails["John"] = "john@gmail.com"
	emails["Doe"] = "doe@gmail.com"
	emails["Foo"] = "foo@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Doe"])

	// Delete from map
	delete(emails, "Foo")
	fmt.Println(emails)

	// Declare map and add key value
	phones := map[string]string{"Samsung": "Note 20 Ultra", "Apple": "Iphone 12 Pro Max"}
	fmt.Println(phones)

	phones["Sony"] = "Experia XZ1"
	fmt.Println(phones)
}
