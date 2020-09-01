package main

import (
	"fmt"
	"math"
)

var funcVar func(int) int

func applyIt(afunct func(int) int, val int) int {
	return afunct(val)
}

func incFn(x int) int { return x + 1 }

func decFn(x int) int { return x - 1 }

// función retorna otra función
func MakeDistOrigin(o_x, o_y float64) func(float64, float64) float64 {
	fn := func(x, y float64) float64 {
		return math.Sqrt(math.Pow(x-o_x, 2) + math.Pow(y-o_y, 2))
	}
	return fn
}

func main() {

	// Usando una función en una variable
	funcVar = incFn
	fmt.Println(funcVar(8))

	// Pasando una función como argumento
	fmt.Println(applyIt(incFn, 2))
	fmt.Println(applyIt(decFn, 2))

	// Función anónima
	v := applyIt(func(x int) int { return x + 1 }, 2)
	fmt.Println(v)

	//Función retorna funcion
	fmt.Println("Origen 1")
	Dist1 := MakeDistOrigin(0, 0)
	fmt.Println("Origen 2")
	Dist2 := MakeDistOrigin(2, 2)
	fmt.Println("Distancia de (2,2) al Origen 1 (0,0)")
	fmt.Println(Dist1(2, 2))
	fmt.Println("Distancia de (2,2) al Origen 2 (2,2)")
	fmt.Println(Dist2(2, 2))

}
