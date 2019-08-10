package main

import "fmt"

func main() {
	// // ARRAYS
	// var fruitArr [2]string

	// // Assign Values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// // Decalre and Assign
	// fruitArr := [2]string{"Apple", "Orange"}

	// SLICING
	fruitSlice := []string{"Apple", "Orange", "Grape"}

	// count the number of items in array use len()
	fmt.Println(len(fruitSlice))

	fmt.Println(fruitSlice)
}
