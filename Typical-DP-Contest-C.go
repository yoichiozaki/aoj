package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())
	r := make([]float64, 1<<uint(k)+1)
	for i := 1; i <= 1<<uint(k); i++ {
		scanner.Scan()
		r[i], _ = strconv.ParseFloat(scanner.Text(), 64)
	}
	// dp[i][j]: iさんがj連勝する確率
	// i = 1, 2, 3, ..., 2^k
	// j = 0, 1, 2, ..., k
	// dp[i][j] =
	//     dp[i][j-1] * dp[k][j-1] * P(iさんがkさんに勝つ確率)
	//     を各kさん（kさんはiさんがjラウンド目で勝負する可能性のある人）について総和を撮ったもの
	dp := make([][]float64, 1<<uint(k)+1)
	for i := range dp {
		dp[i] = make([]float64, k+1)
	}

	// dp[i][0]: iさんが0連勝する確率
	for i := 1; i <= 1<<uint(k); i++ {
		dp[i][0] = 1
	}

	// iさんがj連勝する確率を求めるためにはkさんがj-1連勝する確率が求めらている必要があるのでラウンドごとに計算する
	for j := 1; j <= k; j++ {
		for i := 1; i <= 1<<uint(k); i++ {
			mask1 := 0xffffffff << uint(j-1)
			mask2 := 0xffffffff << uint(j)
			p := 0.0
			for ksan := 1; ksan <= 1<<uint(k); ksan++ {
				if ((i-1)&mask1) != ((ksan-1)&mask1) && ((i-1)&mask2) == ((ksan-1)&mask2) && ksan-1 != i-1 {
					p += dp[ksan][j-1] * (1 / (1 + math.Pow(10, (r[ksan]-r[i])/400)))
				}
			}
			dp[i][j] = dp[i][j-1] * p
		}
	}

	// dp[i][k]: iさんがk連勝する（=優勝する）確率
	for i := 1; i <= 1<<uint(k); i++ {
		fmt.Printf("%.9f\n", dp[i][k])
	}
}
