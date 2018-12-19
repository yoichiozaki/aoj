// n×nの成長度表に現れる数字を重複を許してm-1個選んで掛け合わせる時の最大値を求める
// ->掛け算型のDP

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
	for scanner.Scan() {
		buf := strings.Split(scanner.Text(), " ")
		n, _ := strconv.Atoi(buf[0])
		m, _ := strconv.Atoi(buf[1])
		if n == 0 && m == 0 {
			break
		}
		growth_table := make([][]float64, n)
		for i := range growth_table {
			scanner.Scan()
			buf_ := strings.Split(scanner.Text(), " ")
			tmp := make([]float64, n)
			for j := range tmp {
				tmp[j], _ = strconv.ParseFloat(buf_[j], 64)
			}
			growth_table[i] = tmp
		}
		dp := make([][]float64, m)
		for i := range dp {
			dp[i] = make([]float64, n)
		}
		for i := range dp {
			dp[0][i] = 1
		}
		for k := 1; k < m; k++ {
			// k個使うとき...
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					// 新たにgrowth_table[i][jを使うかどうか
					dp[k][j] = max(dp[k][j], dp[k-1][i]*growth_table[i][j])
					// check(dp)
				}
			}
		}
		ans := 0.0
		for i := 0; i < n; i++ {
			if dp[m-1][i] > ans {
				ans = dp[m-1][i]
			}
		}
		fmt.Printf("%.2f\n", ans)
	}
}

func max(a, b float64) float64 {
	if a < b {
		return b
	}
	return a
}

// func check(dp [][]float64) {
// 	log.Println()
// 	for i := range dp {
// 		log.Println(dp[i])
// 	}
// 	log.Println()
// }
