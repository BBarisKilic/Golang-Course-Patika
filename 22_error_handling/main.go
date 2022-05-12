package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func main() {
	result, err := evenNum(3)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Girdiginiz sayi:", result)
	}

	result2, err2 := sRoot(-3)

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Girdiginiz sayinin karekökü:", result2)
	}

	openMyFile()
}

func evenNum(num int) (int, error) {
	if num%2 != 0 {
		return 0, errors.New("HATA: Cift sayi girmediniz!")
	}

	return num, nil
}

func sRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("HATA: Negatif sayilarin karekökü alinamaz!")
	}

	return math.Sqrt(num), nil
}

func openMyFile() {
	file, err := os.Open("test.txt")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Dosyamiz:", file)
	}
}
