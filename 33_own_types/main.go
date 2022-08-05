package main

import (
	"fmt"
	"strings"
)

type mile float64
type kilometer float64
type mystring string

func main() {
	var m1 mile

	m1 = 3.2

	fmt.Printf("%T, %v", m1, m1)

	f1 := float64(4.4)

	fmt.Println()
	fmt.Println(m1 + mile(f1))
	fmt.Println(float64(m1) + f1)

	k1 := kilometer(7.8)

	fmt.Println()
	fmt.Println(k1)

	k2 := kilometer(4.4)

	fmt.Println(k1 + k2 + 2.1)
	fmt.Printf("%.02f", k1+k2+2.1)
	fmt.Println()

	mString := mystring("arin")

	fmt.Println(strings.ToUpper(string(mString)))
}
