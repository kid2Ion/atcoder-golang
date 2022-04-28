package main

import (
	"bufio"
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

}

/*==========================================
 *             Library
 *==========================================*/

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
