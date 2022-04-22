package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n string
	fmt.Scan(&n)
	arr := strings.Split(n, "")
	a := arr[0]
	b := arr[1]
	c := arr[2]
	abc, _ := strconv.Atoi(a + b + c)
	bca, _ := strconv.Atoi(b + c + a)
	cab, _ := strconv.Atoi(c + a + b)
	fmt.Println(abc + bca + cab)
}
