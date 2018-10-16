package main

import (
	"fmt"
)

func main() {
	fmt.Println(evenSum(sum, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...))
	fmt.Println(oddSum(sum, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...))
}

func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func evenSum(f func(x ...int) int, i ...int) int {
	xe := []int{}
	for _, v := range i {
		if v%2 == 0 {
			xe = append(xe, v)
		}
	}
	return f(xe...)
}

func oddSum(f func(x ...int) int, i ...int) int {
	xe := []int{}
	for _, v := range i {
		if v%2 != 0 {
			xe = append(xe, v)
		}
	}
	return f(xe...)
}
