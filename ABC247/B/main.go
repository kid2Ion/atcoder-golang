package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)

	s := make([]string, n)
	t := make([]string, n)
	names := make(map[string]int)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i], &t[i])
		names[s[i]]++
		if s[i] != t[i] {
			names[t[i]]++
		}
	}

	flg := true
	for i := 0; i < n; i++ {
		if names[s[i]] > 1 && names[t[i]] > 1 {
			flg = false
			break
		}
	}
	if flg {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
