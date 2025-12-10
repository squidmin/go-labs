package main

import "fmt"

func main() {
	str := "Demo: Printing characters"

	// for range
	for pos, char := range str {
		fmt.Printf("Character on position %d: %c\n", pos, char)
	}
	fmt.Println()
	str2 := "Japanese: 日本語"

	// for range
	for pos, char := range str2 {
		fmt.Printf("Character %c starts at byte position %d\n", char, pos)
	}
}
