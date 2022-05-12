package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Task I
	sum, dif, prod := calculation(22, 3)
	fmt.Println("Toplam:", sum)
	fmt.Println("Fark:", dif)
	fmt.Println("Carpim:", prod)

	// Task II
	grade, err := getGrade()
	if err != nil {
		fmt.Printf("HATA: %v\n", err)
	}

	var result string

	if grade >= 50 {
		result = "Gectin"
	} else {
		result = "Kaldin"
	}

	fmt.Println(result)

	// Task III
	runGame()
}

func calculation(num1, num2 int) (int, int, int) {
	sum := num1 + num2
	dif := num1 - num2
	prod := num1 * num2

	return sum, dif, prod
}

func getGrade() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		return 0, err
	}

	num, convertErr := convertToInt(input)

	if convertErr != nil {
		return 0, convertErr
	}

	return num, nil
}

func convertToInt(input string) (int, error) {
	input = strings.TrimSpace(input)

	num, err := strconv.Atoi(input)

	if err != nil {
		return 0, err
	}

	return num, nil
}

func runGame() {
	target := getRandomNumber(1, 100)

	fmt.Println("1 ile 100 arasindaki sayiyi tahmin etmeye calisin!")

	reader := bufio.NewReader(os.Stdin)

	for attempts := 0; attempts < 10; attempts++ {
		fmt.Println(10-attempts, " hakkiniz kaldi.")
		fmt.Print("Lütfen bir sayi giriniz: ")

		input, err := getUserInput(reader)

		if err != nil {
			fmt.Println("HATA:", err)
			continue
		}

		if input > target {
			fmt.Println("Tahmininiz büyük, lütfen daha kücük bir sayi giriniz!")
		} else if input < target {
			fmt.Println("Tahmininiz kücük, lütfen daha büyük bir sayi giriniz!")
		} else {
			fmt.Printf("Tahmininiz dogru! %d. denemede dogru sayiyi buldunuz.\n", attempts+1)
			break
		}
	}
}

func getRandomNumber(min, max int) int {
	rand.Seed(time.Now().Unix())

	return rand.Intn(max-min) + min
}

func getUserInput(reader *bufio.Reader) (int, error) {

	input, err := reader.ReadString('\n')

	if err != nil {
		return 0, err
	}

	num, convertErr := convertToInt(input)

	if convertErr != nil {
		return 0, convertErr
	}

	return num, nil
}
