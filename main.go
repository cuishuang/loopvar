package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// Assuming we want to print 3 numbers in 3 goroutines.
	numbers := []int{1, 2, 3}

	for _, num := range numbers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(num) // The loop variable num is captured
		}()
	}

	wg.Wait()
}
