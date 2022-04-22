package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	arr := strings.Split(s, "")
	n1, _ := strconv.Atoi(arr[0])
	n2, _ := strconv.Atoi(arr[2])
	fmt.Println(n1 * n2)
}
