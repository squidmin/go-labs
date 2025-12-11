// maps_examples.go
package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Age  int
	City string
}

func main() {
	// ---------------------------------------------------------
	// 1. Declaring maps
	// ---------------------------------------------------------

	// A nil map (cannot insert into until initialized)
	var nilMap map[string]int
	fmt.Println("nilMap:", nilMap)

	// Using make
	ages := make(map[string]int)
	ages["Alice"] = 30
	ages["Bob"] = 25

	// Map literal
	fruits := map[string]int{
		"apples":  5,
		"oranges": 3,
		"bananas": 7,
	}

	fmt.Println("ages:", ages)
	fmt.Println("fruits:", fruits)

	// ---------------------------------------------------------
	// 2. Inserting & updating values
	// ---------------------------------------------------------

	ages["Charlie"] = 40          // insert
	ages["Bob"] = ages["Bob"] + 1 // update

	fmt.Println("updated ages:", ages)

	// ---------------------------------------------------------
	// 3. Lookup with comma-ok
	// ---------------------------------------------------------

	value, exists := fruits["apples"]
	fmt.Println("apples:", value, "exists?", exists)

	value, exists = fruits["grapes"]
	fmt.Println("grapes:", value, "exists?", exists)

	// ---------------------------------------------------------
	// 4. Deleting keys
	// ---------------------------------------------------------

	delete(fruits, "oranges")
	fmt.Println("after delete:", fruits)

	// ---------------------------------------------------------
	// 5. Iterating over a map (order NOT guaranteed)
	// ---------------------------------------------------------

	fmt.Println("Iterating over fruits:")
	for key, val := range fruits {
		fmt.Printf("  %s -> %d\n", key, val)
	}

	// ---------------------------------------------------------
	// 6. Ordered iteration using sorted keys
	// ---------------------------------------------------------

	keys := make([]string, 0, len(fruits))
	for k := range fruits {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Println("Ordered iteration:")
	for _, k := range keys {
		fmt.Printf("  %s -> %d\n", k, fruits[k])
	}

	// ---------------------------------------------------------
	// 7. Using structs as map values
	// ---------------------------------------------------------

	people := map[string]Person{
		"Alice":   {Age: 30, City: "Paris"},
		"Bob":     {Age: 26, City: "Tokyo"},
		"Charlie": {Age: 40, City: "New York"},
	}

	fmt.Println("people:", people)
	fmt.Println("Alice lives in:", people["Alice"].City)

	// Updating a field requires reassigning
	p := people["Bob"]
	p.City = "London"
	people["Bob"] = p

	fmt.Println("updated Bob:", people["Bob"])

	// ---------------------------------------------------------
	// 8. Maps of slices
	// ---------------------------------------------------------

	groups := map[string][]string{
		"fruits":     {"apple", "banana"},
		"vegetables": {"carrot", "potato"},
	}

	groups["fruits"] = append(groups["fruits"], "orange")
	fmt.Println("groups:", groups)

	// ---------------------------------------------------------
	// 9. Maps of maps
	// ---------------------------------------------------------

	matrix := map[string]map[string]int{
		"row1": {"col1": 1, "col2": 2},
		"row2": {"col1": 3},
	}

	// Insert into nested map
	matrix["row2"]["col2"] = 4

	// Add a new nested map
	matrix["row3"] = map[string]int{"col1": 5}

	fmt.Println("matrix:", matrix)

	// ---------------------------------------------------------
	// 10. Maps are reference types (copying a map copies pointer)
	// ---------------------------------------------------------

	copyMap := ages
	copyMap["Alice"] = 99

	fmt.Println("ages after modifying copyMap:", ages)
	fmt.Println("copyMap:", copyMap)

	// ---------------------------------------------------------
	// 11. Checking map equality (must compare manually)
	// ---------------------------------------------------------

	a := map[string]int{"x": 1, "y": 2}
	b := map[string]int{"x": 1, "y": 2}

	fmt.Println("a == b?", mapsEqual(a, b))
}

// mapsEqual compares two string-int maps for equality.
func mapsEqual(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for key, valA := range a {
		valB, ok := b[key]
		if !ok || valA != valB {
			return false
		}
	}
	return true
}
