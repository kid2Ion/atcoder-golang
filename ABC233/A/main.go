package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	for i := 0; i <= 100; i++ {
		if (10*i + x) >= y {
			fmt.Println(i)
			break
		}
	}
}
