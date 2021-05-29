This project contains different kind of deadlock code scenario that developer will encounter when writing concurrency code in Golang.

Each folder are structured as follows:

* `problemX` - each folder contains README.md that explain about the concurrency problem and solution
* `problematic` - folder contains the deadlock error code
* `solution` - folder contains the proper solution to resolve the issue. Sometimes it contains multiple solution for a problem

The concurrency problem collected in this repo are from different variety of source from the internet - user groups, slack channel, github, etc

Here are some very good resource to learn more about concurrency in Golang:

* [Golang — Understanding channel, buffer, blocking, deadlock and happy groutines.](https://gist.github.com/YumaInaura/8d52e73dac7dc361745bf568c3c4ba37)
  https://dtyler.io/articles/2021/04/13/sync_cond/
  https://github.com/dty1er/size-limited-queue