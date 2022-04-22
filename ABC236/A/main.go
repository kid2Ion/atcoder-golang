package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	var a, b int
	fmt.Scan(&s, &a, &b)

	arr := strings.Split(s, "")
	result := make([]string, len(arr))
	for i := 0; i < len(arr); i++ {
		if i == a-1 {
			result[i] = arr[b-1]
			continue
		}
		if i == b-1 {
			result[i] = arr[a-1]
			continue
		}
		result[i] = arr[i]
	}
	fmt.Println(strings.Join(result, ""))
}
