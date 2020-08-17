package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Enter text: ")
	var x string
	fmt.Scanln(&x)
	a := strings.Contains(x, "a")
	start := strings.Index(x, "i")
	end := strings.Index(x, "n")
	length := len(x) - 1

	if start == 0 && end == length && a == true {
		print("Found!")
	} else {
		print("Not Found!")
	}

}
