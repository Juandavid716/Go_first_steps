package main

import (
	"fmt"
	"math"
)

// Method has a receiver type that it is associated with this.
// Use dot notation to call the method

// MyInt is used for converting an integer
type MyInt int

// Double convert a integer to double
func (mi MyInt) Double() int {
	return int(mi * 2)
}

//Struct combined with methods

// Point struct
type Point struct {
	x float64
	y float64
}

// DistToOrig calculate the distance from the origin (0,0) to one point
func (p Point) DistToOrig() float64 {
	t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
	return math.Sqrt(t)
}

func main() {
	v := MyInt(3)
	fmt.Println(v.Double())
	p1 := Point{3, 4}
	fmt.Println(p1.DistToOrig())

}
