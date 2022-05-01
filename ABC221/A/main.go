package main

import (
	"bufio"
	"fmt"
	"math"
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
	a, b := nextInt(), nextInt()
	c := pow(32, (a - b))
	fmt.Println(c)
}

/*==========================================
 *             Library/hiro_original
 *==========================================*/

// 一行ずつ読み込む(標準入力)
func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// スペース区切りで読み込む(標準入力)
func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

// 冪乗(intのままでok)
func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
