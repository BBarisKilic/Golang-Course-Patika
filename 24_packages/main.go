package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Packages")

	fmt.Println(strings.Contains("Packages", "ages"))
	fmt.Println(strings.Count("Packages", "a"))
	fmt.Println(strings.ToUpper("Packages"))
}
