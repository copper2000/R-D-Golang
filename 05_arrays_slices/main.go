package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr)

	// Declare and assign
	carArr := [2]string{"Toyota", "Ford"}

	fmt.Println(carArr)

	// Slices
	carSlice := []string{"Honda", "Suzuki", "Chevrolet", "Volvo"}

	fmt.Println(carSlice)
	fmt.Println(len(carSlice))

	fmt.Println(carSlice[0:3])

}
