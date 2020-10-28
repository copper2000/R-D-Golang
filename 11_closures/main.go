package main

import "fmt"

/*
	- What is closure?
	- Closure là những function tham chiếu đến các biến tự do (free avariable) tách biệt.
      Nói cách khác, function được định nghĩa trong closure sẽ ghi nhớ môi trường (environment)
      trong nó được tạo ra.
*/

func adder() func(int) int { // adder is the closure function
	sum := 0
	return func(x int) int { // anonymous func is the function defined in the closure => it will remember the sum
		sum += x
		return sum
	}
}

func main() {
	sum := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}
}
