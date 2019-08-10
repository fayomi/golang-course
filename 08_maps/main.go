package main

import "fmt"

func main() {
	// MAPS are key value pairs

	// Define map
	emails := make(map[string]string)

	// declare map and add KV
	surnames := map[string]string{"Bob": "Jone", "Sharon": "Stone", "Mike": "Hardy"}

	// Assign KV
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	// Delete from Map
	delete(emails, "Bob")

	fmt.Println(emails)
	fmt.Println(surnames)

}
