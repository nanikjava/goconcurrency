The sample code shows how the `context.WithDeadline` works.

In the following code 
```
deadline := time.Now().Add(3 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
```
we are creating `context` with 3secs deadline. The code will execute multiple
goroutine that will accept the `context` as the parameter. Inside the goroutine
it was check when the `context` is _done_ 

Since the `context` has a deadline of 3secs that means any code accepting this context 
will have to complete it's operation within 3secs.

In the example code the following goroutine code

```
go worker(ctx, 8) // execute a job after 8 seconds
```

will not complete it's operation as the `worker` function will execute it's code
after 8secs. This means that the `worker` function will execute the `<-Done()` 
statement as shown below.

```
select {
case <-ctx.Done():
    fmt.Printf("%0.2fs - worker(%ds) killed!\n", time.Since(startTime).Seconds(), seconds)
    return // kills goroutine
....
}
```

The final output of the sample code will be as follows

```
2.00s - worker(2s) completed the job.
3.00s - worker(8s) killed!
3.00s - worker(6s) killed!
Number of active goroutines 1
```
