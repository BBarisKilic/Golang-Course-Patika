package main

import "fmt"

func main() {

	x := 10
	y := 10.0

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	// Type Conversion => type(value) => int(y)

	fmt.Println(x + int(y))

	myFunc()
	mySecondFunc()
	myThirdFunc()
}

func myFunc() {
	var x int8 = 127
	var y int16

	y = int16(x)

	fmt.Println(x, y)
}

func mySecondFunc() {
	x := 10
	y := "10"

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	// fmt.Println(x + int(y))
	// Not possible
}

func myThirdFunc() {
	var x int32 = 106
	y := string(x)

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)
}
