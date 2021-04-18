Problem
---
The `Publish` has a goroutine that has a time expiry limit (in the example code it has 5secs expiry). 

Once the goroutine completes the `<-wait` will throw an error as there are no more goroutine exist that will be using the channel.

Solution
----
The proper way to do this is when the goroutine function inside `Publish` completes it need to close out the channel. 

This is important because the channel was created by `Publish` so it must be responsible to close it out to make sure other parts of the code that 
are using the channel are aware that there are no more data to be processed.