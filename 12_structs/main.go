package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	// Can also be formatted like this
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value reciever) based on struct
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + "years old"
}

// method to change something
// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried method (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "f" {
		p.lastName = spouseLastName
	} else {
		return
	}

}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}

	// Alternative
	person2 := Person{"Bob", "Johnson", "New York", "m", 25}
	fmt.Println(person1)

	// Get a single field
	fmt.Println(person1.firstName)

	// call method to increase age
	person1.hasBirthday()
	person1.getMarried("doe")
	person2.getMarried("Thompson")

	// Call struct method
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}
