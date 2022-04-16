package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	var s float64
	s = 1 / ((a * a) + (b * b))
	x := math.Sqrt(s)
	fmt.Println(a*x, b*x)
}
