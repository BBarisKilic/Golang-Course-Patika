package main

import "fmt"

func main() {
	const x = 5
	const y int8 = 5

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", x+y, x+y)

	myFunc()
	mySecondFunc()
}

func myFunc() {
	const x = int16(5.2 + 4.8)
	const y = 5.2 + 4.8

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
}

func mySecondFunc() {
	const x = 4
	const y = 5.2
	const z = true

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", x+y, x+y)
}
