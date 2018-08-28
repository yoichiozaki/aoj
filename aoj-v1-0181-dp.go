package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

const INF = 0x7fffffff

func main() {
	for scanner.Scan() {
		buf := strings.Split(scanner.Text(), " ")
		m, _ := strconv.Atoi(buf[0])
		n, _ := strconv.Atoi(buf[1])
		if n == 0 && m == 0 {
			break
		}
		dp := make([][][]int, m+1)
		for i := range dp {
			dp[i] = make([][]int, n+1)
			for j := range dp[i] {
				dp[i][j] = make([]int, n+1)
				for k := range dp[i][j] {
					dp[i][j][k] = INF
				}
			}
		}
		k := 0
		w := make([]int, n+1)
		for i := 1; i < n+1; i++ {
			scanner.Scan()
			l, _ := strconv.Atoi(scanner.Text())
			w[i] = w[i-1] + l
			if l > k {
				k = l
			} // k is the maximum width of the books
		}
		if n <= m {
			// If the number of stages in the bookshelf is
			// larger than the number of books,
			// put one book in one stage.
			fmt.Println(k)
			continue
		}
		// use only one stage
		for i := 1; i < n+1; i++ {
			dp[1][1][i] = w[i]
		}
		for k := 2; k < m+1; k++ {
			// use k stages
			for i := 1; i < n+1; i++ {
				for j := i; j < n; j++ {
					for l := j; l < n; l++ {
						if dp[k-1][i][j] < INF {
							dp[k][j+1][l+1] = min(dp[k][j+1][l+1], max(w[l+1]-w[j], dp[k-1][i][j]))
						}
					}
				}
			}
		}
		ans := INF
		for i := 1; i < n+1; i++ {
			if dp[m][i][n] < ans {
				ans = dp[m][i][n]
			}
		}
		fmt.Println(ans)
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
