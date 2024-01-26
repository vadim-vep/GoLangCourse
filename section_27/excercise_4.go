package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	incrementer := 0
	goRoutNum := 100
	wg.Add(goRoutNum)
	var mu sync.Mutex

	for i := 0; i < goRoutNum; i++ {
		go func() {
			mu.Lock()
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Println(incrementer)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("final result", incrementer)
}
