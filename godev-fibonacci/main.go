package main

import (
	"fmt"
)

func main() {
	fmt.Println(fibo(1))
	fmt.Println(fibo(2))
	fmt.Println(fibo(3))
	fmt.Println(fibo(4))
	fmt.Println(fibo(5))
	fmt.Println(fibo(6))
	fmt.Println(fibo(7))
	fmt.Println(fibo(8))
}

func fibo(n int) int {
	if n == 0 {
		return 0
	}
	return n + fibo(n-1)
}
