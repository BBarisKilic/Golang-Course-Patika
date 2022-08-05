package main

import "fmt"

func main() {
	/*	x := 5
		fmt.Println(x)
		double(x)
		fmt.Println(x)*/
	x := 5
	fmt.Println(x)
	double(&x)
	fmt.Println(x)

	arr := []int{1, 10, 100}
	doubleArray(arr)
	fmt.Println(arr)
}

/*func double(x int) {
	x *= 2

	fmt.Println(x)
}*/

func double(x *int) {
	*x *= 2
	fmt.Println(*x)
}

func doubleArray(x []int) {
	for i := 0; i < len(x); i++ {
		x[i] *= 2
	}

	fmt.Println(x)
}
