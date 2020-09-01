package main

import "fmt"

// Función swap
func swap(slc []int, i int) {

	slc[i], slc[i+1] = slc[i+1], slc[i]
}

// Método de ordenamiento Burbuja
func bubbleSort(slc []int) {
	for i := 0; i < len(slc)-1; i++ {
		for j := 0; j < len(slc)-i-1; j++ {
			if slc[j] > slc[j+1] {
				swap(slc, j)
			}
		}

	}

}
func main() {
	var input int
	slc := make([]int, 0, 10)
	fmt.Printf("Enter 10 numbers ")
	for i := 0; i < 10; i++ {
		fmt.Scan(&input)
		slc = append(slc, input)
	}
	fmt.Println("Inicial slice: ", slc)
	bubbleSort(slc)

	fmt.Println("Final slice:", slc)
}
