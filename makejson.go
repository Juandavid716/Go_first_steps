package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// Creación de mapa
	m := make(map[string]string)
	var name string
	var add string
	fmt.Println("Enter name")
	fmt.Scan(&name)
	fmt.Println("Enter address")
	fmt.Scan(&add)

	// Inputs
	m["address"] = name
	m["name"] = add

	//fmt.Println(m)

	// Convertir a JSON
	barr, err := json.Marshal(m)
	fmt.Println(string(barr))
	if err != nil {
		fmt.Println("Error")
	}
}
