package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	// Slow goroutine (3s delay)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "Response from slow goroutine"
	}()

	// Fast goroutine (1s delay)
	go func() {
		time.Sleep(1 * time.Second)
		ch <- "Response from fast goroutine"
	}()

	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout: no response in 2 seconds")
	}
}
