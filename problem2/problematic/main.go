package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	c := make(chan int)
	go readFromChannel(c, time.After(time.Duration(2)*time.Second))
	time.Sleep(time.Duration(5) * time.Second)
	c <- 10
	wg.Wait()
}

func readFromChannel(c chan int, ti <-chan time.Time) {
	select {
	case x := <-c:
		fmt.Println("Read", x)
	case <-ti:
		fmt.Println("TIMED OUT")
	}
	wg.Done()
}
