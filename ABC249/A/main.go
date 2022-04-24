package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, b, c, d, e, f, x, taka, aoki int
	fmt.Fscan(reader, &a, &b, &c, &d, &e, &f, &x)

	c += a
	f += d
	for i := 0; i < x; i++ {
		if i%c < a {
			taka += b
		}
		if i%f < d {
			aoki += e
		}
	}
	if taka > aoki {
		fmt.Println("Takahashi")
	} else if taka == aoki {
		fmt.Println("Draw")
	} else {
		fmt.Println("Aoki")
	}
}
