package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := 1_000_000
	workers := 100 // limit goroutines

	jobChan := make(chan int, workers)
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for job := range jobChan {
				_ = job * 2 // simulate work
			}
		}(i)
	}

	// Send 1 million tasks
	for i := 0; i < jobs; i++ {
		jobChan <- i
	}

	close(jobChan)
	wg.Wait()

	fmt.Println("All tasks completed")
}
