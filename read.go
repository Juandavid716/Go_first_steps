package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {

	var filename string

	fmt.Println("Enter name of file")
	fmt.Scan(&filename)

	// Archivo abierto
	s := make([]Person, 0, 3)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Leyendo archivo
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var cadena = scanner.Text()
		split := strings.Split(cadena, " ")
		v := Person{split[0], split[1]}

		s = append(s, v)
	}
	// Imprimiendo archivo qué contiene un slice de estructuras.
	fmt.Println(s)
	for i := range s {
		fmt.Println("Struct # ", i)
		fmt.Println("name: ", s[i].fname, "lastname: ", s[i].lname)

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
