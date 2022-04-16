package main

import (
	"fmt"
	"math"
)

func main() {
	var s float64
	fmt.Scan(&s)
	fmt.Println(math.Sqrt(s * (s + 12800000)))
}
