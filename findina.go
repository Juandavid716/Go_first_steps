package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// Leyendo cadena
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	// cadena en minusculas
	res1 := strings.ToLower(text)

	// Se verifica si la cadena empieza por i, termina en n y contiene la letra a.
	a := strings.Contains(res1, "a")
	start := strings.Index(res1, "i")
	end := strings.Index(res1, "n")

	// Diferencia de longitudes
	length := len(res1) - 3

	// Resultado
	if start == 0 && end == length && a == true {
		print("Found!")
	} else {
		print("Not Found!")
	}

}
