package main

import (
	"fmt"
)

func main() {
	a := increment()
	b := increment()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func increment() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
