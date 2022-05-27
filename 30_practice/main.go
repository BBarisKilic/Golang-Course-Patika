package main

import "fmt"

func main() {
	myArr := [5]string{"tahran", "ankara", "londra", "berlin", "madrid"}

	for index, value := range myArr {
		fmt.Println(index, value)
	}

	fmt.Println()

	mySlc := []int{1, 2, 3, 4, 5}

	for index, value := range mySlc {
		fmt.Println(index, value)
	}

	mySlc2 := []int{1, 2, 3, 4, 5, 6, 7, 8}

	mySlc3 := mySlc2[:4]

	fmt.Println(mySlc3)

	mySlc4 := mySlc2[4:7]

	fmt.Println(mySlc4)

	mySlc5 := mySlc2[6:]

	fmt.Println(mySlc5)

	mySlc6 := append(mySlc2[2:4], mySlc2[6:8]...)

	fmt.Println(mySlc6)

	myMap := map[string][]string{
		"Yasar Kemal":     {"Ince Memed", "Yer Demir Gök Bakir"},
		"Sabahattin Ali":  {"Kuyucakli Yusuf", "Kürk Mantolu Madonna", "Degirmen"},
		"Haruki Murakami": {"1Q84", "Dans Dans Dans", "Kumandani Öldürmek"},
	}

	fmt.Println(myMap)
	fmt.Println(myMap["Sabahattin Ali"])
	fmt.Println(myMap["Haruki Murakami"][0])

	for key, value := range myMap {
		fmt.Println("Yazarimiz:", key)
		fmt.Println("Bazi kitaplari asagidadir:")
		for i, v := range value {
			fmt.Println("\t", i+1, v)
		}
	}
}
