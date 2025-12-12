// struct_examples.go
package main

import (
	"encoding/json"
	"fmt"
)

// ------------------------------------------------------------
// 1. Basic struct declaration
// ------------------------------------------------------------

type Person struct {
	Name string
	Age  int
}

// ------------------------------------------------------------
// 2. Struct with exported & unexported fields
// ------------------------------------------------------------

type Account struct {
	Owner   string
	Balance float64
	active  bool // unexported field
}

// ------------------------------------------------------------
// 3. Embedded structs (composition)
// ------------------------------------------------------------

type Address struct {
	City  string
	State string
}

type Employee struct {
	Person   // embedded struct
	Position string
	Address  // another embedded struct
}

// ------------------------------------------------------------
// 4. Methods on structs
// ------------------------------------------------------------

// Value receiver (does NOT modify original)
func (p Person) Greet() string {
	return "Hello, my name is " + p.Name
}

// Pointer receiver (can modify original)
func (a *Account) Deposit(amount float64) {
	a.Balance += amount
	a.active = true
}

// ------------------------------------------------------------
// 5. Constructor functions (Go does not have real constructors)
// ------------------------------------------------------------

// NewPerson returns a new Person value.
func NewPerson(name string, age int) Person {
	return Person{Name: name, Age: age}
}

// NewAccount returns a pointer to a new Account.
func NewAccount(owner string) *Account {
	return &Account{Owner: owner, Balance: 0, active: true}
}

// ------------------------------------------------------------
// Main Demonstration
// ------------------------------------------------------------

func main() {

	// ------------------------------------------------------------
	// Basic struct initialization
	// ------------------------------------------------------------

	var p1 Person // zero value
	fmt.Println("p1:", p1)

	p2 := Person{Name: "Alice", Age: 30}
	fmt.Println("p2:", p2)

	p3 := &Person{Name: "Bob", Age: 25} // pointer literal
	fmt.Println("p3:", p3)

	p4 := NewPerson("Charlie", 40) // constructor function
	fmt.Println("p4:", p4)

	// ------------------------------------------------------------
	// Accessing struct fields
	// ------------------------------------------------------------

	fmt.Println("p2 name:", p2.Name)
	p2.Age = 31

	fmt.Println("p3 age before:", p3.Age)
	p3.Age++ // automatic dereference
	fmt.Println("p3 age after:", p3.Age)

	// ------------------------------------------------------------
	// Using methods
	// ------------------------------------------------------------

	fmt.Println(p2.Greet())

	// NOTE: NewAccount returns *Account, not a type.
	acct := NewAccount("Alice")
	acct.Deposit(100)
	fmt.Println("Account:", acct)

	// ------------------------------------------------------------
	// Embedded struct usage
	// ------------------------------------------------------------

	emp := Employee{
		Person:   Person{Name: "Derek", Age: 22},
		Position: "Software Engineer",
		Address:  Address{City: "Seattle", State: "WA"},
	}

	fmt.Println("Employee Name:", emp.Name) // promoted field
	fmt.Println("Employee City:", emp.City) // promoted field

	// ------------------------------------------------------------
	// Anonymous structs
	// ------------------------------------------------------------

	car := struct {
		Make  string
		Model string
		Year  int
	}{
		Make:  "Tesla",
		Model: "Model Y",
		Year:  2024,
	}

	fmt.Println("Anonymous car struct:", car)

	// ------------------------------------------------------------
	// Comparing structs
	// ------------------------------------------------------------

	pA := Person{"Alice", 30}
	pB := Person{"Alice", 30}
	pC := Person{"Bob", 20}

	fmt.Println("pA == pB:", pA == pB)
	fmt.Println("pA == pC:", pA == pC)

	// ------------------------------------------------------------
	// Slice of structs
	// ------------------------------------------------------------

	team := []Person{
		{Name: "Tom", Age: 29},
		{Name: "Jerry", Age: 35},
	}

	for _, member := range team {
		fmt.Println("Team member:", member)
	}

	// ------------------------------------------------------------
	// Map of structs
	// ------------------------------------------------------------

	userMap := map[string]Person{
		"user1": {Name: "Sam", Age: 20},
		"user2": {Name: "Lucy", Age: 28},
	}

	fmt.Println("userMap[user1]:", userMap["user1"])

	// ------------------------------------------------------------
	// JSON marshaling/unmarshaling
	// ------------------------------------------------------------

	jsonData, _ := json.Marshal(p2)
	fmt.Println("JSON:", string(jsonData))

	var decoded Person
	json.Unmarshal(jsonData, &decoded)
	fmt.Println("Decoded JSON:", decoded)
}
