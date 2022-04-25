package main

// GOで四捨五入
import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var n float64
	scanner := bufio.NewReader(os.Stdin)
	fmt.Fscan(scanner, &n)
	fmt.Println(math.Round(n))
}
