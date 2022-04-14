package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n)

	fmt.Scan(&s)
	arr := strings.Split(s, "")
	fmt.Println(arr[n-1])
}
