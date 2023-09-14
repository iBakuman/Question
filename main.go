package main

import (
	"fmt"
	"sync"
)

func resolve(arr []int) {
	fmt.Printf("Before: %v\n", arr)
	input := make(chan int, 100)
	go func() {
		for i := 0; i < len(arr); i++ {
			input <- i
		}
		close(input)
	}()

	var wg sync.WaitGroup
	count := 0
	for idx := range input {
		arr[idx]++
		count++
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				if idx, ok := <-input; ok {
					arr[idx]++
				} else {
					break
				}
			}
		}()
	}
	wg.Wait()
	fmt.Printf("After: %v\n", arr)
	fmt.Printf("len(arr): %d, Created %d goroutines totally!\n", len(arr), count)
}
