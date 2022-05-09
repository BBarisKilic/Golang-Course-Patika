package main

import "fmt"

func main() {
	x, y := 15, 10

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)

	fmt.Println("Toplam:", addition(x, y))
	fmt.Println("Fark:", substruction(x, y))
	fmt.Println("Çarpım:", multiplication(x, y))
	fmt.Println("Bölüm:", division(x, y))
	fmt.Println("Kalan:", remainder(x, y))

	myFunc()
}

func addition(a int, b int) int {
	return a + b
}

func substruction(a int, b int) int {
	return a - b
}

func multiplication(a int, b int) int {
	return a * b
}

// func division(a int, b int) float64 {
// 	return float64(a) / float64(b)
// }

func division(a int, b int) int {
	return a / b
}

func remainder(a int, b int) int {
	return a % b
}

func myFunc() {
	// Increment ++, Decrement --, POSTFIX

	x := 10

	fmt.Println(x)

	x = x + 1

	fmt.Println(x)

	x++

	fmt.Println(x)

	// fmt.Println(x++)
	// Bu satir calismaz cünkü x++ go dilinde bir statement ayni print gibi
	// Ve go dilinde tek satirda iki statement bulunamaz.
}
