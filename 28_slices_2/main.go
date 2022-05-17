package main

import "fmt"

func main() {
	underArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	mySlc := underArray[2:5]
	fmt.Println(mySlc)

	mySlc2 := underArray[:6]
	fmt.Println(mySlc2)

	mySlc3 := underArray[3:]
	fmt.Println(mySlc3)

	mySlc4 := underArray[:]
	fmt.Println(mySlc4)

	mySlc[0] = 100
	fmt.Println(mySlc)
	fmt.Println(mySlc2)
	fmt.Println(mySlc3)
	fmt.Println(mySlc4)

	myFunc()
	myFunc2()
	myFunc3()
	myFunc4()
}

func myFunc() {
	mySlc := []int{1, 2, 3}

	mySlc2 := append(mySlc, 4, 5)
	fmt.Println(mySlc2)

	mySlc2[0] = 100

	fmt.Println(mySlc2)
	fmt.Println(mySlc)
}

func myFunc2() {
	mySlc := []int{1, 2, 3}
	mySlc2 := []int{4, 5}

	mySlc = append(mySlc, mySlc2...)

	fmt.Println(mySlc)
}

func myFunc3() {
	mySlc := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(mySlc)
	// Ilk n elemani kaldirma
	mySlc2 := mySlc[2:]
	fmt.Println(mySlc2)
	// Son n elemani kaldirma
	mySlc3 := mySlc[:len(mySlc)-3]
	fmt.Println(mySlc3)
}

func myFunc4() {
	mySlc := make([]bool, 2)

	fmt.Println(mySlc)

	mySlc2 := make([]float64, 2)

	fmt.Println(mySlc2)

	var mySlc3 []int

	fmt.Printf("%#v\n", mySlc3)

	var mySlc4 = make([]int, 0)

	fmt.Printf("%#v\n", mySlc4)
}
