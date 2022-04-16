package main

import "fmt"

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	if a == 1 {
		if b == 2 || b == 10 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	} else if a == 10 {
		if b == 9 || b == 1 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	} else {
		if b == a+1 || b == a-1 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
