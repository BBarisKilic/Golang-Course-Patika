package main

import "fmt"

type employee struct {
	name      string
	age       int
	isMarried bool
}

type manager struct {
	employee
	hasDegree bool
}

func main() {
	e1 := employee{
		name:      "Baris",
		age:       27,
		isMarried: false,
	}

	e2 := e1

	e2.name = "Oznur"

	fmt.Println(e2.name)
	fmt.Println(e1.name)

	m1 := manager{
		employee: employee{
			name:      "Ulas",
			age:       23,
			isMarried: false,
		},
		hasDegree: true,
	}

	fmt.Printf("%#v\n", m1)

	m2 := manager{}

	m2.name = "Meryem"
	m2.age = 50
	m2.isMarried = false
	m2.hasDegree = false

	fmt.Printf("%#v\n", m2)

	theBoss := struct {
		name  string
		money bool
	}{name: "THE BOSS", money: true}

	fmt.Println(theBoss)
}
