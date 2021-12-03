package main

import "fmt"

func main() {
	var firstName, lastName string = "Barış", "Kılıç"

	var age uint8 = 40
	age = 26

	var isMarried bool
	isMarried = false

	fmt.Println("Name:", firstName, lastName)
	fmt.Println("Age:", age)
	fmt.Println("Is Married?", isMarried)

	fmt.Printf("%T\n", firstName)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isMarried)
}
