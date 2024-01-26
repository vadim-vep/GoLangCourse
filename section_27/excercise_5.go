package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var incrementer int64

	goRoutNum := 100
	wg.Add(goRoutNum)

	for i := 0; i < goRoutNum; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			fmt.Println(incrementer)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("final result", incrementer)
}
