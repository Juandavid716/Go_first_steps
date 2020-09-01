package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	//Creación de slice
	s := make([]int, 0, 3)
	var cont int = 0

	input := bufio.NewScanner(os.Stdin)
	fmt.Print("insert a number until you want. If you want to exit, press x ")
	for input.Scan() {

		if input.Text() == "x" {
			fmt.Println("> exit:", input.Text())
			break
		} else {
			// Agrega números a un slice
			intVar, _ := strconv.Atoi(input.Text())
			fmt.Println(len(s), cont)
			if cont == len(s) {

				s = append(s, intVar)

			} else {

				s[cont] = intVar

			}
			// Slice ordenado
			sort.Ints(s)
			fmt.Println(s)
			cont = cont + 1
		}

	}

}
