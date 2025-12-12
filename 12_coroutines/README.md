# Coroutines

In Go, the "coroutine-like" primitive is called a **coroutine**.
Here's a single, self-contained `.go` file that demonstrates:

- Starting goroutines
- Using `sync.WaitGroup` to wait for them
- Communicating with channels (unbuffered & buffered)
- A simple worker pool
- Using `select` with a timeout

You can name this file `goroutines_examples.go` and run it with:

```bash
go run goroutines_examples.go
```

---

## `goroutines_examples.go`

```go
// goroutines_examples.go
package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// worker simulates doing some work, then sends a result on the channel.
func worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		// Simulate work taking a random time
		sleepMs := rand.IntN(500) + 100 // 100–600 ms
		time.Sleep(time.Duration(sleepMs) * time.Millisecond)

		results <- fmt.Sprintf("Worker %d processed job %d in %dms", id, job, sleepMs)
	}
}
```
