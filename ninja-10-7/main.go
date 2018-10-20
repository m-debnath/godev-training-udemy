package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// send
	send(c)

	// receive
	receive(c)

	fmt.Println("about to exit")
}

func send(c chan<- int) {
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}
		}()
	}
}

func receive(c <-chan int) {
	for i := 0; i < 100; i++ {
		fmt.Println(i, <-c)
	}
}
