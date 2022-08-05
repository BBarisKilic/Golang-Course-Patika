package main

import "fmt"

func main() {
	practice1()
	practice2()
	practice3()
	practice4()
}

type myType int

func practice1() {
	var deger myType
	deger = 10

	fmt.Printf("%T, %v\n", deger, deger)
}

func practice2() {
	x := 10
	y := &x
	*y = 20

	fmt.Println(x)
}

type rectangle struct {
	a, b int
}

func (r rectangle) display() {
	fmt.Println("Dikdörtgenin kenar uzunluklari ", r.a, "ve ", r.b)
}

func (r rectangle) area() int {
	return r.a * r.b
}

func (r rectangle) circumference() int {
	return 2 * (r.a + r.b)
}

func practice3() {
	r1 := rectangle{4, 7}
	r1.display()
	fmt.Println("Dikdörtgenin alani: ", r1.area())
	fmt.Println("Dikdörtgenin cevresi: ", r1.circumference())
}

type family struct {
	name      string
	age       int
	isMarried bool
}

func showFamily() []family {
	family1 := family{
		name:      "Ulas",
		age:       23,
		isMarried: false,
	}

	family2 := family{
		name: "Meryem",
		age:  50,
	}

	family3 := family{
		"Baris",
		27,
		false,
	}

	var family4 family
	family4.name = "Oznur"
	family4.age = 27
	family4.isMarried = false

	return []family{family1, family2, family3, family4}
}

func practice4() {
	families := showFamily()

	for i := 0; i < len(families); i++ {
		fmt.Printf("%v, %T\n", families[i], families[i])
	}
}
