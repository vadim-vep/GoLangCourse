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

	for i := 0; i < goRoutNum; i++ {
		go func() {

			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("final result", incrementer)
}
