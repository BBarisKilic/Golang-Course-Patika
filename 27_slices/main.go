package main

import "fmt"

func main() {
	var myArr = [3]int{1, 2, 3}
	var myArr2 = [...]int{1, 2, 3, 4}

	fmt.Println(myArr, "lenght:", len(myArr))
	fmt.Println(myArr2, "lenght:", len(myArr2))

	var mySlc = []int{1, 2, 3}
	fmt.Println(mySlc, "lenght:", len(mySlc))

	myFunc()
}

func myFunc() {
	var myArr = [4]int{}
	fmt.Println(myArr, "lenght:", len(myArr))
	myArr[0] = 4
	fmt.Println(myArr, "lenght:", len(myArr))

	var mySlc = []int{}
	fmt.Println(mySlc, "lenght:", len(mySlc))
	// mySlc[0] = 4
	// fmt.Println(mySlc, "lenght:", len(mySlc))

	var mySlc2 = make([]int, 4)
	fmt.Println(mySlc2, "lenght:", len(mySlc2))
	mySlc2[0] = 4
	fmt.Println(mySlc2, "lenght:", len(mySlc2))
}
