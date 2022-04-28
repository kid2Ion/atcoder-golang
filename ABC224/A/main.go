package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	initialBufSize = 1e4
	maxBufSize     = 1e6 + 7
)

var buf []byte = make([]byte, initialBufSize)
var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, int(maxBufSize))
}

func main() {
	s := nextLine()
	strs := [2]string{"er", "ist"}
	for _, str := range strs {
		if strings.HasSuffix(s, str) {
			fmt.Println(str)
		}
	}
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}
