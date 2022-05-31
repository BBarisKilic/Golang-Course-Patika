package main

import "fmt"

type employer struct {
	name      string
	age       int
	isMarried bool
}

func main() {

	// 1
	var employee struct {
		name      string
		age       int
		isMarried bool
	}

	fmt.Println(employee)
	fmt.Printf("%#v\n", employee)
	fmt.Println(employee.age)

	employee.name = "Baris"
	employee.age = 27
	employee.isMarried = false

	fmt.Printf("%#v\n", employee)
	fmt.Println(employee.age)

	// 2
	var e1 employer

	e1.name = "Ulas"
	e1.age = 23
	e1.isMarried = false

	var e2 employer

	e2.name = "Oznur"
	e2.age = 27
	e2.isMarried = false

	fmt.Printf("%#v\n", e1)
	fmt.Printf("%#v\n", e2)
}
