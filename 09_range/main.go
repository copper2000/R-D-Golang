package main

import "fmt"

func main() {
	ids := []int{33, 66, 11, 2, 9}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index -> use '_' syntax for unused variable
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	cars := map[string]string{"Toyota": "Land Cruiser", "Ford": "Explorer", "Land Rover": "Range Rove"}

	// loops key and value
	for key, value := range cars {
		fmt.Printf("%s: %s\n", key, value)
	}

	// loops only key
	for key := range cars {
		fmt.Println("Name: " + key)
	}
}
