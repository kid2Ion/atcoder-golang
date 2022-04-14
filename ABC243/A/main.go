package main

import "fmt"

func main() {
	var v, a, b, c int

	fmt.Scan(&v, &a, &b, &c)

	vafter := v % (a + b + c)
	if (vafter - a) < 0 {
		fmt.Println("F")
	} else {
		vafter -= a
		if (vafter - b) < 0 {
			fmt.Println("M")
		} else {
			fmt.Println("T")
		}
	}
}
