package main

import "fmt"

// 排他的論理和（２進数で考えるとわかる）
func main() {
	var a, b, c, d, e, f int

	fmt.Scan(&a, &b, &c, &d, &e, &f)

	fmt.Println(a^c^e, b^d^f)
}
