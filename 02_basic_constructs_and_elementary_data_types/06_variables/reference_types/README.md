# Reference types

Reference types in Go include:
- slices
- maps
- channels
- pointers
- functions

These types refer to data structures that can be modified when passed to functions, unlike value types which are copied.
Demonstrating how reference types work can help in understanding their behavior in a program, especially with regard to mutability and function calls.

---

### Slices and Maps

Below is a simple example demonstrating reference types in Go, specifically focusing on slices and maps, which are commonly used reference types.
This example will show how modifications to these structures inside functions are reflected outside the functions due to their reference nature.

```go
package main

import "fmt"

// Modifies the slice by appending a new element
func modifySlice(slice []int) {
    slice[0] = 10 // Modifying the first element
    slice = append(slice, 20) // Appending a new element
    fmt.Println("Inside modifySlice, slice is now:", slice)
}

// Modifies the map by adding a new key-value pair
func modifyMap(m map[string]int) {
    m["newKey"] = 50
    fmt.Println("Inside modifyMap, map is now:", m)
}

func main() {
    // Demonstrating with a slice
    mySlice := []int{1, 2, 3}
    fmt.Println("Before modifySlice, mySlice is:", mySlice)
    modifySlice(mySlice)
    fmt.Println("After modifySlice, mySlice is:", mySlice)

    // Demonstrating with a map
    myMap := map[string]int{"key1": 10, "key2": 20}
    fmt.Println("Before modifyMap, myMap is:", myMap)
    modifyMap(myMap)
    fmt.Println("After modifyMap, myMap is:", myMap)
}
```

This code demonstrates two functions:
- `modifySlice`, which takes a slice of integers
- `modifyMap`, which takes a map with string keys and integer values.

Both functions modify their respective inputs to show how changes to these reference types inside the functions are visible outside the functions:

In `modifySlice`, we modify the first element of the slice and append a new element to it.
The change to the first element is reflected outside the function, but the append might not reflect if it causes the slice to reallocate.
This is because `append` may return a new slice header if the underlying array needs to be expanded.

In `modifyMap`, we add a new key-value pair to the map.
This modification is reflected outside the function, demonstrating that maps are indeed reference types.
When you run this program, you'll see how the modifications made within the functions affect the original data structures, illustrating the reference type behavior in Go.

---

## Pointers and Channels

Pointers and channels are fundamental to understanding concurrency and memory manipulation in Go.

### Pointers

Pointers hold the memory address of a variable. Unlike some other languages, Go pointers do not support pointer arithmetic for safety reasons.

#### Code

```go
package main

import "fmt"

// Function to modify an integer's value through a pointer
func modifyValue(num *int) {
	*num = 100 // Set the value of the variable at the pointer address to 100
}

func main() {
	originalValue := 10
	fmt.Println("Before modifyValue:", originalValue)

	// Passing the address of originalValue to the function
	modifyValue(&originalValue)

	fmt.Println("After modifyValue:", originalValue)
}
```

In the pointers example, we see how a function can modify the value of a variable outside its own scope by using a pointer to that variable. This demonstrates the ability of pointers to reference and manipulate data directly.

### Channels

Channels are a Go data type that provides a way for goroutines to communicate with each other and synchronize their execution.
Channels are particularly useful for concurrent programming in Go.

#### Code

```go
package main

import (
    "fmt"
    "time"
)

// This function sends a message through a channel
func sendMessage(ch chan<- string, message string) {
    fmt.Println("Sending message:", message)
    ch <- message // Send message to channel
}

func main() {
    messageChannel := make(chan string)

    go sendMessage(messageChannel, "Hello, channels in Go!")

    // Simulate doing something else while the message is being sent
    fmt.Println("Main function doing other work")
    time.Sleep(2 * time.Second) // Wait to ensure the message is sent

    // Receive the message from the channel
    receivedMessage := <-messageChannel
    fmt.Println("Received message:", receivedMessage)
}
```

---

## Functions

Functions in Go can also be treated as reference types when assigned to variables or passed as arguments to other functions.
This allows for higher-order functions, callbacks, and functional programming patterns.

#### Code

```go
package main

import (
	"fmt"
	"strings"
)

// Function type definition
type modifierFunc func(string) string

// Takes a string and returns its uppercase version
func toUpperCase(str string) string {
	return strings.ToUpper(str)
}

// Accepts a string and a function that modifies a string
func printModifiedString(s string, modifier modifierFunc) {
	fmt.Println(modifier(s))
}

func main() {
	originalString := "hello, world"
	fmt.Println("Original string:", originalString)

	// Passing the toUpperCase function as an argument
	printModifiedString(originalString, toUpperCase)
}
```
