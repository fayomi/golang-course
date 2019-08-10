package main

import "fmt"

func main() {
	// MAIN DATA TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	//complex64 complex128

	//CREATING VARIABLES

	// using var
	// var name = "Fayomi"
	var age = 27
	const isCool = true

	// Shorthand declarations but can only be used inside functions
	pet := "dog"
	size := 1.2

	// to assign multiple variables quickly
	name, email := "fayomi", "fayomi@gmail.com"

	fmt.Println(name, age, isCool, pet, size, email)
	// to get the type of value
	fmt.Printf("%T\n", size)

}
