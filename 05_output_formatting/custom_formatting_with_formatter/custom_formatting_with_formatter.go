package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) Format(f fmt.State, c rune) { // Implementing the fmt.Formatter interface
	switch c {
	case 'v': // Using 'v' as a generic verb that could be customized further
		if f.Flag('+') {
			fmt.Fprintf(f, "%s (%d years old)", p.Name, p.Age)
		} else {
			fmt.Fprintf(f, "%s", p.Name)
		}
	}
}

func FormatterExample() {
	person := Person{Name: "Test User", Age: 30}
	fmt.Printf("Name: %v\n", person)    // Default format
	fmt.Printf("Custom: %+v\n", person) // Custom format triggered by '+'
}

func main() {
	FormatterExample()
}
