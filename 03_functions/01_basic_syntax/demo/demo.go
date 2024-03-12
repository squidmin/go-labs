package main

import (
	"fmt"
	"go-labs/03_functions/01_basic_syntax/average"
)

func main() {
	fmt.Println("Result: ", average.average([]float64{98, 93, 77, 82, 83}))
}
