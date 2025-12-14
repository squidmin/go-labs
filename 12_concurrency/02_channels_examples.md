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

func main() {
	// ------------------------------------------------------------
	// 1. Basic unbuffered channel
	// ------------------------------------------------------------
	fmt.Println("=== 1. Unbuffered channel (send/receive) ===")

	ch := make(chan string)

	// Start a goroutine that sends a message
	go func() {
		time.Sleep(300 * time.Millisecond)
		ch <- "hello from goroutine"
	}()

	// This receive blocks until the goroutine sends
	msg := <-ch
	fmt.Println("received:", msg)
	fmt.Println()

	// ------------------------------------------------------------
	// 2. Buffered channel
	// ------------------------------------------------------------
	fmt.Println("=== 2. Buffered channel ===")

	buf := make(chan int, 3) // capacity 3

	// These sends do NOT block until the buffer is full
	buf <- 10
	buf <- 20
	buf <- 30

	fmt.Println("len(buf):", len(buf), "cap(buf):", cap(buf))

	// Receive the values
	fmt.Println(<-buf)
	fmt.Println(<-buf)
	fmt.Println(<-buf)
	fmt.Println()

	// ------------------------------------------------------------
	// 3. Directional channels (send-only, receive-only)
	// ------------------------------------------------------------
	fmt.Println("=== 3. Directional channels ===")

	pings := make(chan string)
	pongs := make(chan string)

	go sendMessage(pings, "ping")
	go pingPong(pings, pongs)

	fmt.Println(<-pongs)
	fmt.Println()

	// ------------------------------------------------------------
	// 4. Closing channels and ranging over them
	// ------------------------------------------------------------
	fmt.Println("=== 4. Closing channels and ranging ===")

	numbers := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
		close(numbers) // signal no more values will be sent
	}()

	// range stops when channel is closed and drained
	for n := range numbers {
		fmt.Println("got:", n)
	}
	fmt.Println()

	// ------------------------------------------------------------
	// 5. Worker pool with channels
	// ------------------------------------------------------------
	fmt.Println("=== 5. Worker pool ===")

	jobs := make(chan int)
	results := make(chan string)

	// Start a few workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	go func() {
		for j := 1; j <= 5; j++ {
			jobs <- j
		}
		close(jobs) // no more jobs
	}()

	// Collect results
	for i := 0; i < 5; i++ {
		fmt.Println(<-results)
	}
	fmt.Println()

	// ------------------------------------------------------------
	// 6. select with timeout and default
	// ------------------------------------------------------------
	fmt.Println("=== 6. select with timeout and default ===")

	// Example A: timeout waiting for a channel
	slowChan := make(chan string)

	go func() {
		time.Sleep(800 * time.Millisecond)
		slowChan <- "finished slow operation"
	}()

	select {
	case v := <-slowChan:
		fmt.Println("received:", v)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("timeout: slow operation took too long")
	}
	fmt.Println()

	// Example B: non-blocking send with default
	nonBlocking := make(chan string)

	select {
	case nonBlocking <- "try send":
		fmt.Println("sent to nonBlocking")
	default:
		fmt.Println("send would block, did not send")
	}

	// Drain if anything was actually sent
	select {
	case v := <-nonBlocking:
		fmt.Println("drained:", v)
	default:
		fmt.Println("nothing to drain")
	}

	fmt.Println("\nChannel demo complete.")
}
```
