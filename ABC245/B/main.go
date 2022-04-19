package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	m := make(map[int]bool)

	for _, ele := range arr {
		m[ele] = true
	}
	var val int

	for i := 0; i <= n; i++ {
		_, ok := m[i]
		if !ok {
			val = i
			break
		}
	}
	fmt.Println(val)
}
