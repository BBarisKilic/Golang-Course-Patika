package main

import "fmt"

func main() {
	for i := 0; i <= 10; i += 5 {
		fmt.Println(i)
	}

	myFunc()
	mySecondFunc()
	myThirdFunc()
	myFourthFunc()
}

func myFunc() {
	i := 0

	for ; i <= 10; i += 5 {
		fmt.Println(i)
	}
}

func mySecondFunc() {

	// for { //Infinite Loop
	// 	fmt.Println("Benim adim Baris")
	// }

	// for i := 0; true; i++ { //Infinite Loop
	// 	fmt.Println(i)
	// }

	i := 10

	for i >= 0 {
		fmt.Println(i)
		i--
	}
}

func myThirdFunc() {
	for i := 0; i <= 10; i++ {
		if i%3 == 0 {
			continue
		}

		fmt.Println(i)
	}
}

func myFourthFunc() {
	for i := 0; i <= 10; i++ {
		if i == 3 {
			break
		}

		fmt.Println(i)
	}
}
