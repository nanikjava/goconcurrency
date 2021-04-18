// Go program to illustrate
// the execution of default case
package main

import "fmt"

// Main function
func main() {

	// Creating a channel
	var c = make(chan int)

	go func() {
		c <- 10
	}()

	select {
	case x1 := <-c:
		fmt.Println("Value: ", x1)
		break
	}

}
