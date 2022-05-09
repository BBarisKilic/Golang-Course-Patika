package main

import "fmt"

func main() {

	// (Print - Println) - Printf

	fmt.Print("Merhaba")
	fmt.Println()
	fmt.Println("Merhaba")
	fmt.Printf("Merhaba")
	fmt.Println()

	myFunc()
	mySecondFunc()
	myThirdFunc()
}

func myFunc() {
	name := "Baris"

	fmt.Print("Benim adim", name)
	fmt.Println()
	fmt.Println("Benim adim", name)
	// fmt.Printf("Benim adim", name)
	fmt.Println()
	fmt.Printf("Benim adim %v", name)
	fmt.Println()
	fmt.Printf("Benim adim %v %T", name, name)
	fmt.Println()
}

func mySecondFunc() {
	x := 100

	// base 2 => %b
	// base 8 => %o
	// base 10 => %d

	fmt.Printf("%b %d %o", x, x, x)
	fmt.Println()
}

func myThirdFunc() {
	name, age := "Baris", 27

	fmt.Print("Benim adim ", name, " ve ben ", age, " yasindayim.")
	fmt.Println()
	fmt.Println("Benim adim", name, "ve ben", age, "yasindayim.")
	fmt.Printf("Benim adim %v ve ben %v yasindayim.", name, age)
}

func myFourthFunc() {
	/*
		var coin string
		var customName string // camel case

		var URL string  // kisaltmalar büyük yazilir
		var HTTP string // kisaltmalar büyük yazilir
	*/
}
