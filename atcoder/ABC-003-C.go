package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Scan()
	buf := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(buf[0])
	k, _ := strconv.Atoi(buf[1])
	scanner.Scan()
	buf_ := strings.Split(scanner.Text(), " ")
	rates := make([]float64, n)
	for i := 0; i < n; i++ {
		rates[i], _ = strconv.ParseFloat(buf_[i], 64)
	}
	ans := 0.0
	sort.Float64s(rates)
	for i := n - k; i < n; i++ {
		ans = (ans + rates[i]) / 2
	}
	fmt.Printf("%f", ans)
}
