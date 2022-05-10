package main

import "fmt"

func main() {

	x := 7

	if x := -5; x < 0 {
		fmt.Println(x, "negatif sayidir.")
	} else if x%2 == 0 {
		fmt.Println(x, "cift sayidir.")
	} else {
		fmt.Println(x, "tek sayidir.")
	}

	fmt.Println(x)

	myFunc()
}

func myFunc() {
	if x := 10; x < 0 {
		fmt.Println(x, "negatif sayidir.")
	} else {
		if x%2 == 0 {
			fmt.Println(x, "cift sayidir.")
		} else {
			fmt.Println(x, "tek sayidir.")
		}
	}
}
