package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	initialBufSize = 1e4
	maxBufSize     = 1e6 + 7
)

var buf []byte = make([]byte, initialBufSize)
var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, maxBufSize)
}

func main() {
	n := nextInt()
	t := nextLine()
	var x, y, direction int
	for i := 0; i < n; i++ {
		switch t[i] {
		case 'S':
			switch direction {
			case 0:
				x++
			case 1:
				y--
			case 2:
				x--
			case 3:
				y++
			}
		case 'R':
			direction++
			direction %= 4
		}
	}
	fmt.Println(x, y)
}

// 一行ずつ読み込む
func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// スペース区切りで読み込む
func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
