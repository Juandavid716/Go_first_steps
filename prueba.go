package main

import (
	"fmt"

	ellipse "github.com/Juandavid716/gohelloworld"
)

func main() {
	// x2 + y2 + Ax + By+ C = 0 and  x2 + y2 + Dx + Ey + F = 0 are circle equation
	e := ellipse.Init{
		9, 2,
	}
	//this will give answer as 0.9749960430435691
	fmt.Println(e.Pythagorean())
	fx := ellipse.GenDisplaceFn(5, 6, 4)
	fmt.Println(fx(3))
}
