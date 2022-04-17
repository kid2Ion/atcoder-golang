package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	var k float64
	fmt.Scan(&a, &b, &k)

	for i := 0.0; i < 100.0; i++ {
		if b <= (a * int(math.Pow(k, i))) {
			fmt.Println(i)
			break
		}
	}
}
