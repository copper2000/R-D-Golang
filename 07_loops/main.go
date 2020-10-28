package main

import "fmt"

func main() {

	// long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("===================")

	// short method
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("===================")

	// FizzBuzz -> this is one of some interview questions
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}
