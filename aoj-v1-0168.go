package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	dp := make([]int, 30)
	dp[0] = 1 // 1段の階段の登り方
	dp[1] = 2 // 2段の階段の登り方
	dp[2] = 4 // 3段の階段の登り方
	for i := 3; i < 30; i++ {
		// n段の階段の登り方は
		// n-1段の階段の登り方+n-2段の階段の登り方+n-3段の階段の登り方
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		fmt.Println(((dp[n-1]+9)/10 + 364) / 365)
	}
}
