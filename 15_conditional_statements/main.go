package main

import "fmt"

func main() {
	x := 27

	if x%2 == 0 {
		fmt.Println(x, "cift sayidir.")
	} else {
		fmt.Println(x, "tek sayidir.")
	}

	myFunc()
}

func myFunc() {
	x := -5

	if x < 0 {
		fmt.Println(x, "negatif sayidir.")
	} else if x%2 == 0 {
		fmt.Println(x, "cift sayidir.")
	} else {
		fmt.Println(x, "tek sayidir.")
	}
}
