package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var mu sync.Mutex

func main() {
	counter := 0
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			counter++
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}
