package main

import "fmt"

func main() {
	// FOR LOOP: Long Method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//FOR LOOP: Short Method
	for x := 10; x <= 20; x++ {
		fmt.Printf("number %d\n", x)
	}

	//FIZZBUZZ
	for p := 1; p <= 100; p++ {
		if p%5 == 0 && p%3 == 0 {
			fmt.Printf("fizzbuzz: %d\n", p)
		} else if p%3 == 0 {
			fmt.Printf("fizz: %d\n", p)
		} else if p%5 == 0 {
			fmt.Printf("buzz: %d\n", p)
		} else {
			fmt.Println(p)
		}
	}

}
