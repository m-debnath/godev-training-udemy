package main

import (
	"fmt"
)

func main() {
	fmt.Println("2 + 3 =", mySum([]int{2, 3}...))
	fmt.Println("42 + 9 + 1 =", mySum([]int{42, 9, 1}...))
	fmt.Println("-1 + 6 =", mySum([]int{-1, 6}...))
	fmt.Println("7 + 4 + 2 + 9 =", mySum([]int{7, 4, 2, 9}...))
}

// mySum returns sum of the input numbers
func mySum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
