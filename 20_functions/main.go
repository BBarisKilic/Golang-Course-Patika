package main

import "fmt"

func main() {
	var x, y, sum int

	x = 10
	y = 5
	sum = x + y

	fmt.Printf("%d ile %d toplami: %d\n", x, y, sum)

	x = 11
	y = 7
	sum = x + y

	fmt.Printf("%d ile %d toplami: %d\n", x, y, sum)

	// moduler programming
	// readable code
	// from complex to simple

	fmt.Printf("%d ile %d toplami: %d\n", x, y, sumFunc(5, 8))
	fmt.Printf("%d ile %d toplami: %d\n", x, y, sumFunc(15, 8))
	fmt.Printf("%d ile %d toplami: %d\n", x, y, sumFunc(5, 82))
	fmt.Printf("%d ile %d toplami: %d\n", x, y, sumFunc(1, 6))

	sumFuncTwo(5, 12)

	hello("Baris", 27)
}

func sumFunc(x, y int) int {
	return x + y
}

func sumFuncTwo(x, y int) {
	fmt.Printf("%d ile %d toplami: %d\n", x, y, x+y)
}

func hello(name string, age int) {
	fmt.Printf("Ad: %v , Yas: %d\n", name, age)
}
