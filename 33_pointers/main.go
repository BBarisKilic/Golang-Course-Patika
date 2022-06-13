package main

import "fmt"

func main() {
	name := "Baris"

	fmt.Println(name)
	fmt.Println(&name) // & --> address operator

	fmt.Println()

	x := 22
	fmt.Println(x)
	fmt.Println(&x)

	fmt.Println()

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", &x, &x)

	fmt.Println()

	var y = &x
	fmt.Printf("%T, %v\n", y, y)

	fmt.Println()

	z := &name
	fmt.Printf("%T, %v\n", z, z)
	fmt.Printf("%T, %v\n", *z, *z)

	x1 := 10
	x2 := &x1

	fmt.Println(x1, x2)
	fmt.Println(x1, *x2)

	*x2 = 3

	fmt.Println(x1, x2)
	fmt.Println(x1, *x2)

	fmt.Println()

	a1 := [4]int{1, 10, 100, 1000}
	a2 := a1

	fmt.Println(a1, a2)

	a2[0] = 3
	fmt.Println(a2)
	fmt.Println(a1)

	fmt.Println()

	b1 := []int{1, 10, 100, 1000}
	b2 := b1

	fmt.Println(b1, b2)

	b2[0] = 3
	fmt.Println(b2)
	fmt.Println(b1)
}
