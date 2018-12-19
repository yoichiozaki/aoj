package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

const (
	MAXN = 1000010
	MOD  = 1e9 + 7
)

func main() {
	scanner.Scan()
	buf := strings.Split(scanner.Text(), " ")
	N, _ := strconv.Atoi(buf[0])
	K, _ := strconv.Atoi(buf[1])
	dp := make([]int, MAXN)
	sum := make([]int, MAXN)
	dp[0] = 1
	sum[1] = 1
	for i := 1; i <= N+1; i++ {
		if i != 1 && i != N {
			dp[i] = (sum[i] - sum[max(i-K, 0)] + MOD) % MOD
		}
		sum[i+1] = (sum[i] + dp[i]) % MOD
	}
	fmt.Println(dp[N+1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
