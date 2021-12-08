package main

import "fmt"

func main() {

	// var (
	// 	firstName string = "Barış"
	// 	lastName  string = "Kılıç"
	// 	age       uint8  = 40
	// 	isMarried bool   = false

	// 	weight float32 = 69.5
	// 	height int16 = 180
	// )

	// var name, age, isMarried, weight, height = "Barış", 26, false, 69.5, 180

	name, age, isMarried, weight, height := "Barış", 26, false, 69.5, 180

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Married?", isMarried)
	fmt.Println("Weight:", weight)
	fmt.Println("Height:", height)
}
