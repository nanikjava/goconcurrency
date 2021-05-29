package main

import (
	"context"
	"fmt"
	"time"
)

// The sample code shows the use of WithDeadline.
// The code will stop with the error 'context deadline exceeded'
// this means that the the select {} statement will not be executed
// as the context has expired.

// If we increase the WithDeadline value d with the following
// 		d := time.Now().Add(50000 * time.Millisecond)
// the code will print out the message
//
//		overslept
func main() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

}