package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Scanln function can be used to take the input from the user in the Golang.

	// Message
	fmt.Println("Enter Your Name: ")

	// var then variable name then variable type
	var name string

	// Taking input from user
	fmt.Scanln(&name)

	// another way to do it
	fmt.Printf("Enter information ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	request := scanner.Text()
}
