package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &s)

	result := "Yes"
	arr := [2]bool{false, false}
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++

		if m[s[i]] > 1 {
			result = "No"
			break
		}
		if 'a' <= s[i] && s[i] <= 'z' {
			arr[0] = true
		} else {
			arr[1] = true
		}
	}

	if !arr[0] || !arr[1] {
		result = "No"
	}
	fmt.Println(result)
}
