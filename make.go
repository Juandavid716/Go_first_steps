package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var address string
	fmt.Printf("Enter the name: ")
	fmt.Scanln(&name)
	fmt.Printf("Enter Address: ")
	fmt.Scanln(&address)

	my := map[string]string{
		"Your name is":    name,
		"Your address is": address,
	}
	printMap(my)
	barr, error := json.Marshal(my)

	println("Json object contain:", barr)
	println("error contain", error)
}
func printMap(c map[string]string) {
	for name, address := range c {
		fmt.Println(name)
		fmt.Println(address)

	}
}
