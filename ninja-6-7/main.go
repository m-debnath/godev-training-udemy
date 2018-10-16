package main

import (
	"fmt"
)

func main() {
	x := func() {
		for i := 0; i <= 100; i++ {
			fmt.Println(i)
		}
	}
	x()
}
