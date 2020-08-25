package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	i, _ := strconv.Atoi("10")
	y := i * 2
	fmt.Println(y)

	s := strings.Replace("ianianian", "ni", "in", 2)
	fmt.Println(s)
}
