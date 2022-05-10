package main

import "fmt"

func main() {
	x, y := 3, 7

	fmt.Printf("%T, %v\n", x == y, x == y)
	fmt.Printf("%T, %v\n", x < y, x < y)
	fmt.Printf("%T, %v\n", x > y, x > y)
	fmt.Printf("%T, %v\n", x != y, x != y)

	myFunc()
}

func myFunc() {
	x, y := 15, 20

	set1 := x == y
	set2 := x < y

	fmt.Printf("%T, %v\n", set1 && set2, set1 && set2)

	fmt.Printf("%T, %v\n", set1 || set2, set1 || set2)
}
