Problem
---

The `c` channel has a read timeout of 2 seconds which will expired way before the `c` channel is populated using the `c <- 10`.

When the `readFromChannel` completes golang identifies that there are no more goroutines in existence that is waiting for the `c` channel, so it throws
the deadlock error.

Solution
----

The code in `solution1` increase the timeout (to 6 seconds) for `readFromChannel` so that when the 5 seconds sleep expired the code that reads the `c` channel is still 
in existence.

The code in `solution2` does not have timeout for `readFromChannel`. The function will directly call `wg.Done()` once the value is received from the channel. 
This way there is no dependencies on 2 different timeouts.