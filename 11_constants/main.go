package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(areaOfCircle(3.0))

	myFunc()
}

func areaOfCircle(r float64) float64 {
	const pi float64 = 3.14

	return pi * math.Pow(r, 2)
}

func myFunc() {
	const x = 2

	// 	x = 5
	// 	x++
	// 	x = x - 1

	// const y = math.Pow(3, 2)
	// runtime da hesaplanacak deger
	// bundan dolayi compile time da atanan degere yazilamaz

	// z := 4
	// const a = z
	// variable runtimeda hesaplanir ve atanir
	// bu yüzden yukaridaki kod calismaz

	fmt.Printf("%T %v", x, x)
}
