package main

import (
	"fmt"
	"strconv"
)

// func main() {

// 	x := 75
// 	var y float64

// 	y = float64(x)

// 	fmt.Println(y)
// }

// func main() {

// 	x := 5
// 	y := 10

// 	fmt.Println("X:", x, "Y:", y)

// 	x, y = y, x

// 	fmt.Println("X:", x, "Y:", y)
// }

// func main() {
// 	yaş := 27

// 	fmt.Println("Benim yaşım:", yaş)
// }

// func main() {
// 	x := 5

// 	if true {
// 		x := 10
// 		x++
// 		fmt.Println(x)
// 	}

// 	fmt.Println(x)
// }

func main() {
	x := 65
	y := string(int32(x))

	fmt.Printf("%v %T\n", x, x) // 65
	fmt.Printf("%v %T\n", y, y) // A

	y = strconv.Itoa(x)

	fmt.Printf("%v %T\n", y, y) // "65"
}
