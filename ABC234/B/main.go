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
	n := nextInt()
	x := make([]int, n+1)
	y := make([]int, n+1)
	for i := 1; i <= n; i++ {
		x[i] = nextInt()
		y[i] = nextInt()
	}
	var ans int
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			ans = max(ans, (x[i]-x[j])*(x[i]-x[j])+(y[i]-y[j])*(y[i]-y[j]))
		}
	}
	fmt.Println(math.Sqrt(float64(ans)))
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
