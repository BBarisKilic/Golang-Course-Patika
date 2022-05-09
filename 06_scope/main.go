package main

import "fmt"

var packVar = "Package Scope"
var funcVar = "Func(Package) Scope"

func main() {

	if true {
		var blockVar = "Block Scope"
		fmt.Println(blockVar)
	}

	funcVar := "Func Scope"

	fmt.Println(funcVar)

	fmt.Println(packVar)

	myFunc()

	mySecondFunc()
}

func myFunc() {
	fmt.Println(funcVar)

	fmt.Println(packVar)
}

func mySecondFunc() {
	name := "Bülent"
	name, surname := "Baris", "Kilic"
	println(name, surname)
}
