package main

import (
	"fmt"

	"github.com/m-debnath/godev-training-udemy/test-benchmark/03-example/acdc"
)

func main() {
	fmt.Println(acdc.Sum(2, 3))
	fmt.Println(acdc.Sum(-1, 1))
	fmt.Println(acdc.Sum(4, 3))
}
