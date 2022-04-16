package main

import (
	"fmt"
	"math"
)

func main() {
	var s float64
	fmt.Scan(&s)
	if math.Pow(2, s) > s*s {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
