package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker1 %d starting\n", id)

	time.Sleep(time.Duration(2) * time.Second)
	fmt.Printf("Worker1 %d done\n", id)
}
func worker2(id int) {
	fmt.Printf("Worker2 %d starting\n", id)

	time.Sleep(time.Duration(2) * time.Second)
	fmt.Printf("Worker2 %d done\n", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(2)

		iter := i

		go func() {

			worker(iter)
			wg.Done()
		}()

		go func() {

			worker2(iter)
			wg.Done()
		}()

	}

	wg.Wait()

}
