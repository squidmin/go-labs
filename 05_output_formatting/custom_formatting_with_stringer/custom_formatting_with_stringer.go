package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string { // Implementing the fmt.Stringer interface
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

func StringerExample() {
	person := Person{Name: "Test User", Age: 30}
	fmt.Println(person)
}

func main() {
	StringerExample()
}
