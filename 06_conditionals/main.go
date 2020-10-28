package main

import "fmt"

func main() {
	x := 15
	y := 10

	if x <= y {
		fmt.Printf("%d is less than or equal to %d", x, y)
	} else {
		fmt.Printf("%d is less than %d \n", y, x)
	}

	// else if

	color := "blue"

	if color == "blue" {
		fmt.Println("color is blue")
	} else if color == "red" {
		fmt.Println("color is red")
	} else {
		fmt.Println("color is not blue or red")
	}

	// Switch

	switch color {
	case "blue":
		fmt.Println("color is blue")
	case "red":
		fmt.Println("color is red")
	default:
		fmt.Println("color is not blue or red")
	}
}
