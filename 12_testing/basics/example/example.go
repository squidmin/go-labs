package main

import (
	"errors"
)

var errorMessage = "width and height must be positive"

// CalcArea Calculate the area of a rectangle
func CalcArea(w, h int) (int, error) {
	if w < 1 || h < 1 {
		return 0, errors.New(errorMessage)
	}
	return w * h, nil
}

func main() {
	// Example: 3x4 rectangle
	area, err := CalcArea(3, 4)
	if err != nil {
		println(err)
	} else {
		println(area)
	}

	// Example: 0x4 rectangle
	area, err = CalcArea(0, 4)
	if err != nil {
		println(err)
	} else {
		println(area)
	}
}
