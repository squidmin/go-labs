package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

func main() {
	p1 := Point{1, 2}
	p2 := p1
	p2.X = 10
	fmt.Println(p1, p2)
}
