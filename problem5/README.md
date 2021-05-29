Problem
---
The `main.go` is throwing error as follows

```
sleep 1
sleep 2
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
        /home/nanik/Downloads/temp/packages/src/github.com/nanikjava/src/concurrency/problem5/problematic/main.go:18 +0x73

Process finished with the exit code 2

```

Solution
----
The deadlock issue is because the `ch1` channel is waiting to receive from the channel `case <-ch1:` but there is there no 
code that is sending or closing the channel.

The solution is to either send or close the `ch1` channel like following:

```
go func() {
    fmt.Println("sleep 1")
    time.Sleep(3 * time.Second)
    fmt.Println("sleep 2")
    close(ch1)
}()
```