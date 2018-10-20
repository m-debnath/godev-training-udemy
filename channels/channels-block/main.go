package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 1)
	go func() {
		c <- 42
		fmt.Println("In goroutine")
		time.Sleep(time.Second * 10)
	}()
	fmt.Println(<-c)
}
