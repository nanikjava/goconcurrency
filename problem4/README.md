Problem
---

From the error

```
goroutine 10 [chan send]:
main.readFromWorker(0xc00002c240, 0xc000014130)
        /home/nanik/Downloads/temp/packages/src/github.com/nanikjava/src/concurrency/problem4/problematic/main.go:41 +0x174
created by main.main
        /home/nanik/Downloads/temp/packages/src/github.com/nanikjava/src/concurrency/problem4/problematic/main.go:77 +0x151
```

It is seen that the error occured inside `readFromWorker` function, specifically in the following code
```
stageIn1 <- v
```

what does this means ?

this means that the channel that `v` is being read `inCh` is closed. The for loop has no idea why it is closed.

The deadlock also means that there is no goroutine exist when the `main` function closed the channel (it's called `resultsCh` in main).  

Solution
----
Before the `readFromWorker` was like the following

```
func readFromWorker(inCh <-chan *data, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	stageIn1 := make(chan *data)
	stageOut1 := make(chan *data)

	for v := range inCh {
		fmt.Println("v", v)

		stageIn1 <- v
	}
	fmt.Println("outside for v := range inCh ")

	go stage1(stageIn1, stageOut1)
	go stage2(stageOut1)
}
```

and the solution is to placed goroutine _before_ doing `for{}` loop reading the channel.

```
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
```

also we need to place a timeout before closing everything to allow enough time for data to be completed processed in memory. This is done in
the `main` function as follows:

```
...
...
	time.Sleep(2 * time.Second)
	close(inputsCh)
...
...
```