package main

import "fmt"

// Modifies the slice by appending a new element
func modifySlice(slice []int) {
	slice[0] = 10             // Modifying the first element
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
