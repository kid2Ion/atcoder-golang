package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, k, a int

	scanner := bufio.NewReader(os.Stdin)
	fmt.Fscan(scanner, &n, &k, &a)

	r := (k + a - 1) % n
	if r == 0 {
		fmt.Println(n)
	} else {
		fmt.Println(r)
	}
}
