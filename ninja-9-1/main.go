package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("Goroutine:", runtime.NumGoroutine())
	wg.Add(2)
	go foo(1)
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("Goroutine:", runtime.NumGoroutine())
	go foo(2)
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("Goroutine:", runtime.NumGoroutine())
	bar()
	wg.Wait()
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("Goroutine:", runtime.NumGoroutine())
}

func foo(inp int) {
	fmt.Println("In foo", inp)
	wg.Done()
}

func bar() {
	runtime.Gosched()
	fmt.Println("In bar")
}
