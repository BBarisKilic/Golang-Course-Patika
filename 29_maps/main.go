package main

import "fmt"

func main() {
	myMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}

	fmt.Println(myMap, "lenght:", len(myMap))

	myMap["e"] = 5

	fmt.Println(myMap, "lenght:", len(myMap))

	delete(myMap, "a")

	fmt.Println(myMap, "lenght:", len(myMap))

	studentsGrade := make(map[string]int)

	studentsGrade["John"] = 5
	studentsGrade["Mary"] = 6
	studentsGrade["Peter"] = 7

	fmt.Println(studentsGrade, "lenght:", len(studentsGrade))

	_, ok := studentsGrade["John"]

	fmt.Println(ok)

	_, ok2 := studentsGrade["Baris"]

	fmt.Println(ok2)

	sg := studentsGrade

	delete(studentsGrade, "John")

	fmt.Println(sg, "lenght:", len(sg))
	fmt.Println(studentsGrade, "lenght:", len(studentsGrade))

	sg["Firat"] = 8

	fmt.Println(sg, "lenght:", len(sg))
	fmt.Println(studentsGrade, "lenght:", len(studentsGrade))

	for k, v := range sg {
		fmt.Println(k, v)
	}
}
