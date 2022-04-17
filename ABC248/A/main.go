package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	arr := strings.Split(s, "")
	all := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, v := range arr {
		all = remove(all, v)
	}
	fmt.Println(all[0])
}

func remove(s []string, search string) []string {
	result := []string{}
	for _, v := range s {
		if v != search {
			result = append(result, v)
		}
	}
	return result
}
