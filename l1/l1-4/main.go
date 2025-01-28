package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const NWorkers = 5

// Work — implementation of worker pool that processes tasks from a channel.
func Work(ctx context.Context, workerID int64, jobs <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case data, ok := <-jobs:
			if !ok {
				// Channel is closed, exit the worker.
				fmt.Printf("Worker %d: channel closed, shutting down\n", workerID)
				return
			}
			// Print the data to stdout.
			fmt.Printf("Worker %d: %s\n", workerID, data)
		case <-ctx.Done():
			// Context is cancelled, exit the worker.
			fmt.Printf("Worker %d: received shutdown signal\n", workerID)
			return
		}
	}
}

func main() {
	// SIGINT — interrupt signal such as Ctrl+C,
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT)

	// Create a channel for data,
	dataChan := make(chan struct{})

	// Create a context for graceful shutdown,
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// WaitGroup to wait for all workers to finish,
	wg := &sync.WaitGroup{}

	// Spawn workers
	for workerID := int64(1); workerID <= NWorkers; workerID++ {
		wg.Add(1)
		go Work(ctx, workerID, dataChan, wg)
	}

	// Write dummy data,
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // Context is cancelled, stop writing to the channel,
			default:
				// Write data to the channel,
				dataChan <- struct{}{}
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// Lock gracefull shutdown lock,
	<-shutdown

	// Cancel the context to signal workers to stop
	fmt.Println("\nReceived shutdown signal. Stopping workers...")
	cancel()

	// Close the data channel to signal workers that no more data will be sent.
	close(dataChan)

	// Wait for all workers to finish.
	wg.Wait()
	fmt.Println("All workers have shut down gracefully. 🍷🗿")
}
