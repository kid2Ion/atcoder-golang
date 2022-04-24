package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	if n >= 42 {
		fmt.Println("AGC0" + strconv.Itoa(n+1))
	} else if n >= 10 {
		fmt.Println("AGC0" + strconv.Itoa(n))
	} else {
		fmt.Println("AGC00" + strconv.Itoa(n))
	}
}
