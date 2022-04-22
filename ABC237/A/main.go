package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Scan(&n)
	m := math.Pow(2, 31)
	if -m <= n && n < m {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
