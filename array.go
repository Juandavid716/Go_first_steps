package main

import "fmt"

func main() {
	// Arrays

	var x [5]int
	x[0] = 2
	fmt.Print(x)

	// Arrays with default size

	y := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(y)

	// for loops

	for i, v := range y {
		fmt.Println("")
		fmt.Printf("ind %d, val %d", i, v)
	}
	//Slices
	var s []int = y[1:4]
	fmt.Println(s)

	//other slices

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
