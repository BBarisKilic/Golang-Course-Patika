package main

import "fmt"

func main() {
	taskOne()
	taskTwo()
	taskThree()
	taskFour()
	taskFive()
}

func taskOne() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "cift sayidir.")
		} else {
			fmt.Println(i, "tek sayidir.")
		}
	}
}

func taskTwo() {
	x := 10

	for x > 0 {
		fmt.Println(x)
		x--
	}
}

func taskThree() {
	switch x := 25; {
	case x < 20:
		fmt.Printf("%d sayisi 20'den kücük\n", x)
		fallthrough
	case x < 50:
		fmt.Printf("%d sayisi 50'den kücük\n", x)
		fallthrough
	case x < 100:
		fmt.Printf("%d sayisi 100'den kücük\n", x)
		fallthrough
	case x < 200:
		fmt.Printf("%d sayisi 200'den kücük\n", x)
	}
}

func taskFour() {
	x := 20

	if x%2 == 0 {
		fmt.Println(x, "sayisi cifttir.")
		return
	}

	fmt.Println(x, "sayisi tektir.")
}

func taskFive() {
	var x, y int

	for x = 2; x < 50; x++ {
		for y = 2; y < (x / y); y++ {
			if x%y == 0 {
				break
			}
		}

		if y > (x / y) {
			fmt.Printf("%d bir asal sayidir.\n", x)
		}
	}
}
