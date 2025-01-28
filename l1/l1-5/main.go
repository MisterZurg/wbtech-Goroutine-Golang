package main

import (
	"fmt"
	"time"
)

var NSeconds = 2 * time.Second

// Pipline?
func main() {
	pipe := make(chan int)

	timer := time.NewTimer(NSeconds)
	defer timer.Stop()

	// Write to pipe
	go func() {
		cnt := 0
		for {
			fmt.Println("Sending:", cnt)
			pipe <- cnt
			cnt++
			time.Sleep(250 * time.Millisecond)
		}
	}()

	go func() {
		for data := range pipe {
			fmt.Printf("Received: %d\n", data)
			time.Sleep(250 * time.Millisecond)
		}
		fmt.Println("Channel closed, exiting...")
	}()

	for {
		select {
		case <-timer.C:
			fmt.Println("The time has CUM, closing pipe")
			close(pipe)
			return
		}
	}
}
