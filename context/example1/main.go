package main

// Sample code to demonstrate the usage of  Context.cancel
import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// channel to send square of integers
var c = make(chan int)

// send square of numbers
func square(ctx context.Context) {
	i := 0

	for {
		select {
		case <-ctx.Done():
			return // kill goroutine
		case c <- i * i:
			i++
		}
	}
}

// main goroutine
func main() {

	// create cancellable context
	ctx, cancel := context.WithCancel(context.Background())

	go square(ctx) // start square goroutine

	fmt.Println("Number of active goroutines", runtime.NumGoroutine())


	// get 5 square
	for i := 0; i < 5; i++ {
		fmt.Println("Next square is", <-c)
	}

	// cancel context
	cancel() // instead of `defer context()`

	// do other job
	time.Sleep(3 * time.Second)

	// print active goroutines
	fmt.Println("Number of active goroutines", runtime.NumGoroutine())
}
