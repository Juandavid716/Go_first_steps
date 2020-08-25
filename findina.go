package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	res1 := strings.ToLower(text)

	a := strings.Contains(res1, "a")

	start := strings.Index(res1, "i")
	end := strings.Index(res1, "n")

	length := len(res1) - 3

	if start == 0 && end == length && a == true {
		print("Found!")
	} else {
		print("Not Found!")
	}

}
