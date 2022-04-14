package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if a < c {
		fmt.Println("Takahashi")
	} else if a > c {
		fmt.Println("Aoki")
	} else if a == c {
		if b < d {
			fmt.Println("Takahashi")
		} else if b > d {
			fmt.Println("Aoki")
		} else if b == d {
			fmt.Println("Takahashi")
		}
	}
}
