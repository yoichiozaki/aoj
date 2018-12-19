package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	p := make([]int, n+1)
	buf := strings.Split(scanner.Text(), " ")
	for i := range buf {
		p[i+1], _ = strconv.Atoi(buf[i])
	}
	MAX := 0
	for i := range p {
		MAX += p[i]
	}
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, MAX+1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j <= MAX; j++ {
			if p[i] <= j {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-p[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	cnt := 0
	for j := 0; j <= MAX; j++ {
		if dp[n][j] {
			cnt++
		}
	}
	fmt.Println(cnt)
}
