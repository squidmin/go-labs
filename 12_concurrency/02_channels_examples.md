# Channels

The below example `.go` file demonstrates key channel patterns in Go:

- Unbuffered channels (send / receive, blocking)
- Buffered channels
- Channel directions (`chan<-`, `<-chan`)
- Closing channels & ranging over them
- `select` with timeout and `default`
- Using channels to coordinate goroutines

```go
// channels_examples.go
package main

import (
	"fmt"
	"time"
)

// sendMessage demonstrates a send-only channel parameter.
func sendMessage(ch chan<- string, msg string) {
	ch <- msg
}

// pingPong demonstrates using directional channels for input / output
func pingPong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- "pong to: " + msg
}

// worker reads jobs from a channel and sends results on another.
func worker(id int, jobs <-chan int, results chan<- string) {
	for job := range jobs {
		// Simulate work
		time.Sleep(200 * time.Millisecond)
		results <- fmt.Sprintf("worker %d processed job %d", id, job)
    }
}

func main() {}
```
