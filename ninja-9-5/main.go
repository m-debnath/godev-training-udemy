package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	var counter uint64
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddUint64(&counter, 1)
			runtime.Gosched()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}
