package main

import "fmt"

func main() {
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Scan(&arr[i])
	}
	result1 := arr[0]
	result2 := arr[result1]
	result3 := arr[result2]
	fmt.Println(result3)
}
