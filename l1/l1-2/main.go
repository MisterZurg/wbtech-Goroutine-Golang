package main

import (
	"fmt"
	"sync"
)

var array = []int64{2, 4, 6, 8, 10}

func square(wg *sync.WaitGroup, n int64) int64 {
	defer wg.Done()
	return n * n
}

func main() {
	wg := &sync.WaitGroup{}

	for _, n := range array {
		wg.Add(1)
		go func() {
			fmt.Printf("%d^2 = %d\n", n, square(wg, n))
		}()
	}

	wg.Wait()
}
