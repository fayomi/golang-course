package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// use to read val from memory address
	fmt.Println(*b)

	// change value with pointer
	*b = 10
	fmt.Println(a, b)
}
