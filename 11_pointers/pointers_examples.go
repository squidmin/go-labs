// pointers_examples.go
package main

import "fmt"

// Person ------------------------------------------------------------
// Struct for demonstration
// ------------------------------------------------------------
type Person struct {
	Name string
	Age  int
}

// Function using value semantics
func incrementValue(x int) {
	x++
}

// Function using pointer semantics
func incrementPointer(x *int) {
	*x++
}

// Function that returns a pointer (safe in Go)
func newCounter() *int {
	v := 100
	return &v
}

// Modify struct through pointer
func birthday(p *Person) {
	p.Age++
}

func main() {

	// ------------------------------------------------------------
	// 1. Basic pointer declaration
	// ------------------------------------------------------------
	var a int = 10
	var p *int = &a // p stores address of a

	fmt.Println("a:", a)
	fmt.Println("Address of a (&a):", &a)
	fmt.Println("p (address stored):", p)
	fmt.Println("Value pointed to (*p):", *p)

	// Modify through pointer
	*p = 20
	fmt.Println("\nUpdated a via pointer:", a)

	// ------------------------------------------------------------
	// 2. Using * and & operators
	// ------------------------------------------------------------
	b := 5
	ptrB := &b

	fmt.Println("\nb before:", b)
	*ptrB = 15
	fmt.Println("b after (*ptrB = 15):", b)

	// ------------------------------------------------------------
	// 3. Pointers vs values in functions
	// ------------------------------------------------------------
	x := 10
	incrementValue(x) // does NOT change x
	fmt.Println("\nx after incrementValue:", x)

	incrementPointer(&x) // DOES change x
	fmt.Println("x after incrementPointer (&x):", x)

	// ------------------------------------------------------------
	// 4. Nil pointers
	// ------------------------------------------------------------
	var n *int
	fmt.Println("\nnil pointer n:", n)

	// Check for nil before use
	if n == nil {
		fmt.Println("n is nil, assigning new memory")
		val := 99
		n = &val
	}
	fmt.Println("n now points to:", *n)

	// ------------------------------------------------------------
	// 5. Pointers to structs
	// ------------------------------------------------------------
	p1 := Person{Name: "Alice", Age: 29}
	p2 := &p1 // pointer to struct

	fmt.Println("\nStruct p1:", p1)
	fmt.Println("Pointer p2:", p2)

	// Automatic dereferencing when accessing fields
	p2.Age = 30
	fmt.Println("Updated p1 via p2:", p1)

	birthday(&p1)
	fmt.Println("After birthday:", p1)

	// ------------------------------------------------------------
	// 6. Returning pointers from functions
	// ------------------------------------------------------------
	counter := newCounter()
	fmt.Println("\nCounter returned from function:", *counter)

	// ------------------------------------------------------------
	// 7. Comparing pointers
	// ------------------------------------------------------------
	a1 := 5
	a2 := 5

	ptr1 := &a1
	ptr2 := &a1
	ptr3 := &a2

	fmt.Println("\nptr == ptr2:", ptr1 == ptr2) // true (same variable)
	fmt.Println("ptr1 == ptr3", ptr1 == ptr3)   // false (different variables)

	// ------------------------------------------------------------
	// 8. Special case: slices and maps behave like pointers
	// ------------------------------------------------------------
	s := []int{1, 2, 3}
	sCopy := s // not a deep copy

	sCopy[0] = 99
	fmt.Println("\nSlices share underlying array:")
	fmt.Println("s:", s)
	fmt.Println("sCopy:", sCopy)

	m := map[string]int{"x": 1}
	mCopy := m
	mCopy["x"] = 42

	fmt.Println("\nMaps share underlying storage:")
	fmt.Println("m:", m)
	fmt.Println("mCopy:", mCopy)
}
