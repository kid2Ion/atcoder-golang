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
	l, r := nextInt(), nextInt()
	s := next()
	var ans string
	ans += s[:l-1]
	ans += reverse(s[l-1 : r])
	ans += s[r:]
	fmt.Println(ans)
}

/*==========================================
 *             Library_origin
 *==========================================*/

// 一行読み込み
func next() string {
	sc.Scan()
	return sc.Text()
}

// intで読み込み(1文字)
func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

// floatで読み込み(1文字)
func nextFloat64() float64 {
	f, _ := strconv.ParseFloat(next(), 64)
	return f
}

// 空白区切り→配列(int)
func nextIntList(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = nextInt()
	}
	return arr
}

// 空白区切り→配列(float64)
func nextFloatList(n int) []float64 {
	arr := make([]float64, n)
	for i := 0; i < n; i++ {
		arr[i] = nextFloat64()
	}
	return arr
}

// 空白区切り→配列(string)
func nextStringList(n int) []string {
	arr := make([]string, n)
	for i := 0; i < n; i++ {
		arr[i] = next()
	}
	return arr
}

// 冪乗(intのままでok)(xのy乗)
func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

// 比較して大きい数字返却
func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

// 比較して小さい数字返却
func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

// stringのリバースを返す(abcd→dcba)
func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

/*==========================================
 *             Library_default_package
 *==========================================*/

// 辞書順比較
// s < t (string同士の比較)

// 文字列→配列
// arr := strings.Split(s, "")

// 配列→文字列
// s := strings.Join(arr, "")

// sort(int）(返り値なし)
// sort.Ints(arr)

// sort(アルファベット順)(返り値なし)
// sort.Strings(arr)

// 排他的論理和(XOR)(aとbの)
// a ^ b

// 平方根
// math.Sqrt(float64(x))

// string(英数字)の切り取り
// s[1:3]→1文字目から3文字目まで
// s[:2]→2文字目まで
