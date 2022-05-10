package main

import "fmt"

func main() {
	grade := 3

	switch grade {
	case 5:
		fmt.Println("Pekiyi")

	case 4:
		fmt.Println("Iyi")

	case 3:
		fmt.Println("Orta")

	case 2:
		fmt.Println("Gecer")

	case 1:
		fmt.Println("Basarisiz")

	default:
		fmt.Println("Gecersiz Not")
	}

	if grade == 5 {
		fmt.Println("Pekiyi")
	} else if grade == 4 {
		fmt.Println("Iyi")
	} else if grade == 3 {
		fmt.Println("Orta")
	} else if grade == 2 {
		fmt.Println("Gecer")
	} else if grade == 1 {
		fmt.Println("Basarisiz")
	} else {
		fmt.Println("Gecersiz Not")
	}

	myFunc()
	mySecondFunc()
}

func myFunc() {

	switch grade := -3; grade {
	case 5:
		fmt.Println("Pekiyi")

	case 4:
		fmt.Println("Iyi")

	case 3:
		fmt.Println("Orta")

	case 2:
		fmt.Println("Gecer")

	case 1:
		fmt.Println("Basarisiz")

	default:
		fmt.Println("Gecersiz Not")
	}
}

func mySecondFunc() {

	switch grade := 4; grade {
	case 5:
		fmt.Println("Pekiyi")

	case 4:
		fmt.Println("Iyi")
		y := 80
		fmt.Println(y)

	case 3:
		fmt.Println("Orta")

	case 2:
		fmt.Println("Gecer")

	case 1:
		fmt.Println("Basarisiz")

	default:
		fmt.Println("Gecersiz Not")
	}
}
