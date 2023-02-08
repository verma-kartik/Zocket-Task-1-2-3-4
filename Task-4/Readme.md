# Task-4

## Fetching web resources concurrently

To show how Go addresses concurrency, I decided to build a program which would concurrently fetch various web resources, wait for all of them to be fetched, then process them all at once

The ouptput looks like ...

```
go run main.go                                                                              
Fetching https://golang.org/                                                                  
Fetching https://kubernetes.io/                                                               
Fetching https://www.rubyconf.org/ 
Fetching https://www.microsoft.com/en-in 
........https://kubernetes.io/ was fetched
.https://www.microsoft.com/en-in was fetched
......https://golang.org/ was fetched
....https://www.rubyconf.org/ was fetched
```

As you can see from the print statements, the 4 urls are triggered in a sequential way, but the responses come back in different orders due to different server latencies and response transfer time.
