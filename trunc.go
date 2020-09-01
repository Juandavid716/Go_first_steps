package main

import "fmt"

func trun() {

	var i float32
	var y int16
	fmt.Println("Enter float number ")
	fmt.Scan(&i)
	// Float a Integer
	y = int16(i)
	fmt.Println("Number converted to int: ", y)
}
