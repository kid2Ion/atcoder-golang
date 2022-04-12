package main

import (
	"fmt"
	"strings"
)

func main() {
	var n string
	fmt.Scanf("%s", &n)
	arr := strings.Split(n, "")
	result := make([]string, 4)
	result[0] = "0"
	result[1] = arr[0]
	result[2] = arr[1]
	result[3] = arr[2]
	fmt.Println(strings.Join(result, ""))
}
