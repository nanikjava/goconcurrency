package main

import (
	"fmt"
	"sync"
	"time"
)

type data struct {
	id    int
	url   string
	field int
}

type job struct {
	id  int
	url string
}

func sendToWorker(id int, inCh <-chan job, outCh chan<- *data, wgInput *sync.WaitGroup) {
	defer func() {
		wgInput.Done()
	}()

	for v := range inCh {
		// some pre process stuff and then pass to pipeline
		outCh <- &data{id: v.id, url: v.url}
		fmt.Println("outCh " + v.url)
	}
}

func readFromWorker(inCh <-chan *data, wgResult *sync.WaitGroup) {
	defer func() {
		wgResult.Done()
	}()

	stageIn1 := make(chan *data)
	stageOut1 := make(chan *data)

	go stage1(stageIn1, stageOut1)
	go stage2(stageOut1)

	for v := range inCh {
		fmt.Println("v", v)
		stageIn1 <- v
	}

}

func stage1(in chan *data, out chan *data) {
	for s := range in {
		fmt.Println("stage1 = ", s)
		out <- s
	}
}

func stage2(in chan *data) {
	for s := range in {
		fmt.Println("stage2 = ", s)
	}
}

func main() {
	const chanBuffer = 10

	inputsCh := make(chan job, chanBuffer)
	resultsCh := make(chan *data, chanBuffer)

	wgInput := &sync.WaitGroup{}
	wgResult := &sync.WaitGroup{}

	for i := 1; i <= 4; i++ {
		wgInput.Add(1)
		go sendToWorker(i, inputsCh, resultsCh, wgInput)
	}

	wgResult.Add(1)
	go readFromWorker(resultsCh, wgResult)

	for j := 1; j <= 10; j++ {
		inputsCh <- job{id: j, url: "google.com"}
	}

	time.Sleep(2 * time.Second)
	close(inputsCh)
	wgInput.Wait()
	close(resultsCh)
	wgResult.Wait()
}
