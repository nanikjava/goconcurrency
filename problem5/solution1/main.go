package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	go func() {
		fmt.Println("sleep 1")
		time.Sleep(3 * time.Second)
		fmt.Println("sleep 2")
		close(ch1)
	}()

	for {
		select {
		case <-ch1:
			fmt.Println("ch1 pop one")
			return
		}
	}
}
