package main

import "fmt"

func main() {
	// 1
	cities1 := [4]string{"Istanbul", "Tahran", "Berlin", "Roma"}

	fmt.Println(cities1)

	// 2
	cities2 := [4]string{"London", "Pekin", "Moskova", "Atina"}

	fmt.Println(cities2[0])
	fmt.Println(cities2[3])

	fmt.Println(len(cities2))

	// 3
	var myArray [4]int

	fmt.Println(myArray)

	myArray[len(myArray)-1] = 100

	fmt.Println(myArray)

	// 4
	var myArr [4]int
	var myArr2 [4]int

	myArr[2] = 3

	if myArr == myArr2 {
		fmt.Println("Equal!")
	} else {
		fmt.Println("Not equal!")
	}

	// 5
	cities := [4]string{"Istanbul", "Tahran", "Berlin", "Roma"}

	for i := 0; i < len(cities); i++ {
		fmt.Println(i, cities[i])
	}

	cities[0] = "Ankara"

	for i := 0; i < len(cities); i++ {
		fmt.Println(i, cities[i])
	}

	// 6
	myNormalArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	mySquareArray := mySquare(myNormalArray)

	fmt.Println(mySquareArray)

	// 7
	for index, city := range cities {
		fmt.Println(index, city)
	}
}

func mySquare(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
	}

	return arr
}
