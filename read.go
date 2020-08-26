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

	s := make([]Person, 0, 3)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var cadena = scanner.Text()
		split := strings.Split(cadena, " ")
		v := Person{split[0], split[1]}

		s = append(s, v)
	}
	fmt.Println(s)
	for i := range s {
		fmt.Println("Struct # ", i)
		fmt.Println("name: ", s[i].fname, "lastname: ", s[i].lname)

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
