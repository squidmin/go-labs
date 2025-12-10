# Functions and types

## `basics_syntax_types.go` – Basic Syntax and Types

```go
// basics_syntax_types.go
package main

import "fmt"

// Person demonstrates a simple struct type.
type Person struct {
	Name string
	Age int
}

func main() {
	// ----- Basic types -----
	var age int = 30           // int
	height := 1.75             // float64
	var name string = "Alice"  // string
	var isMember bool          // bool, zero value = false
	
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Is member:", isMember)
	
	// Zero values
	var zeroInt int
	var zeroString string
	var zeroBool bool
	fmt.Println("Zero values:", zeroInt, zeroString, zeroBool)

	// ----- Arrays (fixed size) -----
	var scores [3]int = [3]int{10, 20, 30}
	primes := [...]int{2, 3, 5, 7, 11}  // size inferred
	
	fmt.Println("Scores:", scores)
	fmt.Println("Primes:", primes)

	// ----- Slices (dynamic size) -----
	nums := []int{1, 2, 3}
	nums = append(nums, 4, 5)
	fmt.Println("Nums slice:", nums)
	
	// Slicing operation (like nums[1:3])
	sub := nums[1:3]
	fmt.Println("Sub-slice:", sub)

	// ----- Maps (key-value) -----
	fruits := map[string]int{
		"apples": 5,
		"oranges": 10,
    }
	
	// Add / update an entry
	fruits["bananas"] = 7
	
	fmt.Println("Fruits map:", fruits)
	fmt.Println("Number of apples:", fruits["apples"])
	
	// Check if key exists
	value, ok := fruits["pears"]
	fmt.Println("Pears:", value, "exists?", ok)

	// ----- Structs -----
	person := Person{
		Name: "Bob",
		Age: 25,
    }
	
	fmt.Println("Person struct:", person)
	fmt.Println("Person name:", person.Name)
}
```
