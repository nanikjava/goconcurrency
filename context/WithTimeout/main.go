package main

import (
	"context"
	"fmt"
	"time"
)

// The sample code works the same with the WithDeadline package sample
func main() {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), 1200 * time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1150 * time.Millisecond):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}

}