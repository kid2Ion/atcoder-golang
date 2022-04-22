package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	f1 := t*t + 2*t + 3
	f2 := f1 + t
	f3 := f2*f2 + 2*f2 + 3
	f4 := f1*f1 + 2*f1 + 3
	f5 := (f3+f4)*(f3+f4) + 2*(f3+f4) + 3
	fmt.Println(f5)
}
