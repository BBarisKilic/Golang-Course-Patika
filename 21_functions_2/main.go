package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	hello("Baris", 27)

	fmt.Println(result(55))

	myTime()

	input()

	bölüm, kalan := bölme(111, 4)

	fmt.Printf("111 ile 4'ün bölümü %d, kalan ise %d\n", bölüm, kalan)
}

func hello(name string, age int) {
	fmt.Printf("Adim %s, yasim %d\n", name, age)
}

func result(grade int) string {
	if grade >= 50 {
		return "gectiniz"
	}
	return "kaldiniz"
}

func myTime() {
	var moment time.Time = time.Now()

	fmt.Println(moment)
}

func input() {
	fmt.Print("Lütfen Sinav Sonucunuzu Giriniz: ")

	reader := bufio.NewReader(os.Stdin)

	value, _ := reader.ReadString('\n') // _ blank identifier

	fmt.Println(value)
}

func bölme(bölünen, bölen int) (bölüm, kalan int) {
	bölüm = bölünen / bölen
	kalan = bölünen % bölen

	return bölüm, kalan
}
