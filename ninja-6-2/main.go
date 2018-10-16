package main

import (
	"fmt"
)

func main() {
	fmt.Println(foo([]int{1, 2, 1, 2}...))
	fmt.Println(bar([]int{1, 2, 1, 2}))
}

func foo(n ...int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	return (foo(x...))
}
