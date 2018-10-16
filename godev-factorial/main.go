package main

import (
	"fmt"
)

func main() {
	fmt.Println(factorial(1))
	fmt.Println(factorial(5))
	//fmt.Println(factorial(63))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
