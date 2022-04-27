package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &s)
	sarr := strings.Split(s, "")
	count := make(map[string]int)
	for i := 0; i < len(sarr); i++ {
		count[sarr[i]]++
	}
	if len(count) == 1 {
		fmt.Println(1)
	} else if len(count) == 2 {
		fmt.Println(3)
	} else {
		fmt.Println(6)
	}
}
