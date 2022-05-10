package main

import "fmt"

// func main() {
// 	x := 50

// 	// x = x - 10	assignment statement

// 	// x -= 10	assignment operation

// 	x /= 10

// 	fmt.Printf("%T, %v\n", x, x)
// }

// func main() {
// 	F := -40

// 	K := float64(F-32)/1.8 + 273

// 	fmt.Printf("%T, %v\n", K, K)
// }

// func main() {
// 	const myAge = 40

// 	fmt.Printf("%T, %v\n", myAge, myAge)
// }

// const x = 14

// func main() {
// 	const x = 40

// 	fmt.Printf("%T, %v\n", x, x)
// }

// func main() {
// 	const x = 40  // typless
// 	const y = 4.6 // typless

// 	fmt.Printf("%T, %v\n", x+y, x+y)
// }

func main() {
	const x float64 = 6.4
	y := 4 + x

	fmt.Printf("%T, %v\n", y, y)
}
