package main

import (
	"fmt"
)

func main() {
	x := 42.0
	y := 6.0
	fmt.Println(subtract(x, y))
}

// subtract returns difference between two float64 numbers
func subtract(x, y float64) float64 {
	return x - y
}
