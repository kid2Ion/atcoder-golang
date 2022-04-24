package main

import "fmt"

func main() {
	var s, t, x int
	fmt.Scan(&s, &t, &x)
	if s < t {
		if s <= x && x < t {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	} else {
		if s <= x && x <= 23 {
			fmt.Println("Yes")
		} else if 0 <= x && x < t {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
