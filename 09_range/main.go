package main

import "fmt"

func main() {

	ids := []int{33, 23, 54, 67, 65}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("sum: ", sum)

	// range with maps
	surnames := map[string]string{"Bob": "Jone", "Sharon": "Stone", "Mike": "Hardy"}

	for k, v := range surnames {
		fmt.Printf("%s:%s\n", k, v)
	}

}
