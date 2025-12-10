// standard_library_examples.go
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func main() {
	// ----- fmt: formatted I/O -----
	name := "Gopher"
	age := 5
	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)

	// ----- time: working with dates and times -----
	now := time.Now()
	fmt.Println("Current time:", now)

	// Formatting time
	fmt.Println("Formatted:", now.Format(time.RFC3339))

	// Adding durations
	tomorrow := now.Add(24 * time.Hour)
	fmt.Println("Tomorrow:", tomorrow)

	// ----- strings: string manipulation -----
	text := "Go is an expressive, concise, clean, and efficient language."
	containsGo := strings.Contains(text, "Go")
	upper := strings.ToUpper(text)
	replaced := strings.ReplaceAll(text, "Go", "Golang")

	fmt.Println("Original:", text)
	fmt.Println("Contains 'Go'?", containsGo)
	fmt.Println("Uppercase:", upper)
	fmt.Println("Replaced:", replaced)

	// ----- math/rand: random numbers -----
	// Seed the random generator
	rand.Seed(time.Now().UnixNano())

	randomInt := rand.Intn(100)         // 0 <= n < 100
	randomFloat := rand.Float64()       // 0.0 <= f < 1.0
	randomIntRange := rand.Intn(10) + 1 // 1..10
	fmt.Println("Random int (0-99):", randomInt)
	fmt.Println("Random float (0-1):", randomFloat)
	fmt.Println("Random int (1-10):", randomIntRange)

	// ----- sort: sorting slices -----
	ints := []int{5, 3, 9, 1, 7}
	strs := []string{"banana", "apple", "cherry"}

	fmt.Println("Before sort (ints):", ints)
	sort.Ints(ints)
	fmt.Println("After sort (ints):", ints)

	fmt.Println("Before sort (strs):", strs)
	sort.Strings(strs)
	fmt.Println("After sort (strs):", strs)
}
