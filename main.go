// How can you handle 1 million goroutines in Go? worker_pool.go
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	const Tasks = 1_000_000
	const Workers = 1000 // tune to your CPU cores × 2–4

	tasks := make(chan int, 10000)
	var wg sync.WaitGroup

	// Start fixed worker pool
	for i := 0; i < Workers; i++ {
		go func() {
			for task := range tasks {
				// Simulate work
				time.Sleep(1 * time.Millisecond)
				_ = task * task
				wg.Done()
			}
		}()
	}

	fmt.Printf("Sending %d tasks to %d workers...\n", Tasks, Workers)

	start := time.Now()
	for i := 0; i < Tasks; i++ {
		wg.Add(1)
		tasks <- i
	}
	close(tasks)
	wg.Wait()

	fmt.Printf("Done in %.2f seconds\n", time.Since(start).Seconds())
	fmt.Printf("Goroutines: %d\n", runtime.NumGoroutine())
}
