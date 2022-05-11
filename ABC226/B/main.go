package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
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
	sc.Buffer(buf, maxBufSize)
}

func main() {
	n := nextInt()
	m := make(map[string]int, 0)
	for i := 0; i < n; i++ {
		a := nextInt()
		arr := nextStringList(a)
		mkey := strings.Join(arr, ",")
		m[mkey]++
	}
	fmt.Println(len(m))
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

// "Yes","No"の出力
func PrintBool(a bool) {
	if a {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

// 最小数字を返す(引数は複数) ex, nmin(1, 3, 6)
func nmin(a ...int) int {
	ret := a[0]
	for _, v := range a {
		ret = min(ret, v)
	}
	return ret
}

// 最大数字を返す(引数は複数) ex, nmax(1, 3, 6)
func nmax(a ...int) int {
	ret := a[0]
	for _, v := range a {
		ret = max(ret, v)
	}
	return ret
}

// スライスから特定の要素削除(int)
func removei(ints []int, search int) []int {
	result := []int{}
	for _, v := range ints {
		if v != search {
			result = append(result, v)
		}
	}
	return result
}

// スライスから特定の要素削除(string)
func removes(strings []string, search string) []string {
	result := []string{}
	for _, v := range strings {
		if v != search {
			result = append(result, v)
		}
	}
	return result
}

// sliceに特定の要素の存在確認(各型→bool, err)
func contains(target interface{}, list interface{}) (bool, error) {

	switch list.(type) {
	default:
		return false, fmt.Errorf("%v is an unsupported type", reflect.TypeOf(list))
	case []int:
		revert := list.([]int)
		for _, r := range revert {
			if target == r {
				return true, nil
			}
		}
		return false, nil

	case []float64:
		revert := list.([]float64)
		for _, r := range revert {
			if target == r {
				return true, nil
			}
		}
		return false, nil

	case []string:
		revert := list.([]string)
		for _, r := range revert {
			if target == r {
				return true, nil
			}
		}
		return false, nil
	}

	return false, fmt.Errorf("processing failed")
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

// string(s)の3文字目
// s[2]

// s→tは何文字後か(アルファベット)(zの次はa)
// (s + 26 - t) % 26

// sは特定の文字列(t)を含むかどうかを判断
// strings.Contains(s, t)

// 同一文字列を繰り返す
// s := strings.Repeat("s", 10)
